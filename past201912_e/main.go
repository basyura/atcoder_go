package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scanf("%d %d", &n, &cnt)
	relation := newMap(n)

	i := 0
	for i < cnt {
		o := 0
		fmt.Scanf("%d", &o)
		// follow
		if o == 1 {
			var from, to int
			fmt.Scanf("%d %d", &from, &to)
			relation.Follow(from, to)
		} else if o == 2 { // follow back
			from := 0
			fmt.Scanf("%d", &from)
			relation.FollowBack(from)
		} else if o == 3 { // follow follow
			from := 0
			fmt.Scanf("%d", &from)
			relation.FollowFollow(from)
		}
		i++
	}

	i = 1
	for i <= n {
		fmt.Println(relation.Status(i))
		i++
	}
}

type Relations map[int]bool

type RelationStore struct {
	count int
	store map[int]Relations
}

func (self *RelationStore) GetRelations(uid int) Relations {
	return self.store[uid]
}

func (self *RelationStore) Follow(from, to int) {
	relation := self.store[from]
	relation[to] = true
}

func (self *RelationStore) FollowBack(from int) {
	for k, r := range self.store {
		if v, ok := r[from]; ok && v {
			self.store[from][k] = true
		}
	}
}

// フォローフォロー: その時点でユーザ a がフォローしている各ユーザ x に対し、ユーザ a が次を行う
// ユーザ x がフォローしているすべてのユーザ (ユーザ a 自身を除く) をフォローする
func (self *RelationStore) FollowFollow(from int) {
	// from's follow
	relations := self.GetRelations(from)
	targets := []int{}
	for uid := range relations {
		// follow's follow
		for to := range self.GetRelations(uid) {
			if to != from {
				targets = append(targets, to)
			}
		}
	}

	for _, v := range targets {
		self.Follow(from, v)
	}
}

func (self *RelationStore) Status(from int) string {
	follows := self.GetRelations(from)
	s := ""
	i := 1
	for i <= self.count {
		if _, ok := follows[i]; ok {
			s += "Y"
		} else {
			s += "N"
		}
		i++
	}

	return s
}

func newMap(n int) *RelationStore {
	store := map[int]Relations{}
	i := 1
	for i <= n {
		store[i] = Relations{}
		i++
	}
	relation := &RelationStore{
		count: n,
		store: store,
	}
	return relation
}
