package main

import "fmt"

// в голанге нет типа "множество" как set в питоне. можно сымитировать через тип Map.
// но конкретно для этой задачи более эффективный алгоритм - сортировка слайса.

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	mp := map[string]struct{}{}
	var res []string
	for _, v := range input {
		mp[v] = struct{}{}
	}
	for k, _ := range mp {
		res = append(res, k)
	}
	fmt.Println(res)
}
