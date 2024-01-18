package main

import "fmt"

func main() {
	set1 := []int{1, 2, 5, 9, 7, 3, 4}
	set2 := []int{1, 3, 5, 8, 7, 0, 11}
	var res []int
	// Проверка схождения через map[]
	mp := make(map[int]bool)
	for _, v := range set1 {
		mp[v] = true
	}
	for _, v := range set2 {
		if mp[v] == true {
			res = append(res, v)
		}
	}
	fmt.Println(res)
}
