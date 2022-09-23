package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	max := 0
	n--
	for n >= 1 {
		k := 0
		for k < n {
			v := 0
			fmt.Scanf("%d", &v)
			if v > 0 {
				max += v
			}
			k++
		}
		n--
	}
	fmt.Println(max)
}
