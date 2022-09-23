package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	s := ""
	fmt.Scanf("%s", &s)

	reg := regexp.MustCompile(`\d*`)
	if reg.MatchString(s) {
		if v, err := strconv.Atoi(s); err == nil {
			fmt.Println(v * 2)
			return
		}
	}
	fmt.Println("error")
}
