package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	dic := map[int]bool{}

	i := 0
	for i < n {
		var d int
		fmt.Scanf("%d", &d)
		if _, ok := dic[d]; ok {
			dic[d] = true // dup
		} else {
			dic[d] = false
		}
		i++
	}

	dup := -1
	none := -1
	i = 0
	for i < n {
		i++
		v, ok := dic[i]
		if !ok {
			none = i
			continue
		}
		if v {
			dup = i
		}
	}

	if dup == -1 {
		fmt.Println("Correct")
	} else {
		fmt.Printf("%d %d\n", dup, none)
	}
}
