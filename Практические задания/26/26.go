package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s1 := "abcde"
	s2 := "aabcd"
	s3 := "abcdee"

	fmt.Println(uniqueCheck(s1))
	fmt.Println(uniqueCheck(s2))
	fmt.Println(uniqueCheck(s3))
}

func uniqueCheck(str string) bool {
	r := []rune(strings.ToLower(str))
	slices.Sort(r)

	for i := 1; i < len(r); i++ {
		if r[i] == r[i-1] {
			return false
		}
	}
	return true
}
