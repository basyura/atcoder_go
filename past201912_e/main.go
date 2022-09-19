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

type Relation struct {
	count int
	dic   map[int]map[int]bool
}

func (self *Relation) Follow(from, to int) {
	self.dic[from][to] = true
}

func (self *Relation) FollowBack(from int) {
	for k, r := range self.dic {
		if v, ok := r[from]; ok && v {
			self.dic[from][k] = true
		}
	}
}

// フォローフォロー: その時点でユーザ a がフォローしている各ユーザ x に対し、ユーザ a が次を行う
// ユーザ x がフォローしているすべてのユーザ (ユーザ a 自身を除く) をフォローする
func (self *Relation) FollowFollow(from int) {
	// from's follow
	dic := self.dic[from]
	targets := []int{}
	for uid := range dic {
		// follow's follow
		for to := range self.dic[uid] {
			if to != from {
				targets = append(targets, to)
			}
		}
	}

	for _, v := range targets {
		self.Follow(from, v)
	}
}

func (self *Relation) Status(from int) string {
	follows := self.dic[from]
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

func newMap(n int) *Relation {
	dic := map[int]map[int]bool{}
	i := 1
	for i <= n {
		dic[i] = map[int]bool{}
		i++
	}
	relation := &Relation{
		count: n,
		dic:   dic,
	}
	return relation
}
