package main

import "fmt"

func main() {
	// Исходный срез
	nums := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	index := 3

	// Удаляем элемент из среза
	nums = append(nums[:index], nums[index+1:]...)

	fmt.Println(nums)
}
