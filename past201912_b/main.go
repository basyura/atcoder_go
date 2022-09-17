package main

import (
	"fmt"
)

func main() {

	var cnt int
	fmt.Scanf("%d", &cnt)
	values := []int{}

	i := 0
	for i < cnt {
		var v int
		fmt.Scanf("%d", &v)
		values = append(values, v)
		i++
	}

	i = 1
	for i < cnt {
		pre := values[i-1]
		cur := values[i]
		if pre < cur {
			fmt.Println("up", cur-pre)
		} else if pre == cur {
			fmt.Println("stay")
		} else {
			fmt.Println("down", pre-cur)
		}
		i++
	}
}
