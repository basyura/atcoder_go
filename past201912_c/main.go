package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	ints := []int{}

	i := 0
	for i < 6 {
		var s string
		fmt.Scan(&s)
		ints = append(ints, to_i(s))
		i++
	}

	sort.Ints(ints)
	fmt.Println(ints[len(ints)-3])

}

func to_i(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
