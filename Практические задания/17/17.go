package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	target := 9
	if index, err := searchIndex(arr, target); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(index)
	}
}

func searchIndex(sortedArr []int, element int) (int, error) {
	left := 0
	right := len(sortedArr) - 1
	res := binarySearch(sortedArr, element, left, right)
	if res == -1 {
		return -1, errors.New("Данный элемент отсутствует в массиве")
	}
	return res, nil
}

func binarySearch(arr []int, element int, left int, right int) int {
	if right <= left {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == element {
		return mid
	} else if element < arr[mid] {
		return binarySearch(arr, element, left, mid)
	} else {
		mid++
		return binarySearch(arr, element, mid, right)
	}
}
