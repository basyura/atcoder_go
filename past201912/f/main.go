// https://atcoder.jp/contests/past201912-open/tasks/past201912_f
package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	line := ""
	fmt.Scanf("%s", &line)

	reg := regexp.MustCompile(`[A-Z]`)
	list := []string{}

	word := ""
	for _, r := range line {
		s := string(r)
		if !reg.MatchString(s) {
			word += s
			continue
		}

		if word != "" {
			list = append(list, word+s)
			word = ""
		} else {
			word += s
		}
	}

	sort.Slice(list, func(i, j int) bool {
		si := strings.ToLower(list[i])
		sj := strings.ToLower(list[j])
		v := strings.Compare(si, sj)
		if v > 0 {
			return false
		}
		return true
	})

	fmt.Println(strings.Join(list, ""))
}
