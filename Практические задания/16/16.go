package main

import (
	"fmt"
	"math/rand"
)

type QuickSort struct {
	arr []int
}

func main() {
	arr := make([]int, 0)

	for i := 0; i != 200; i++ {
		randomNumber := rand.Intn(1000) + 1
		arr = append(arr, randomNumber)
	}
	num := QuickSort{
		arr: append(arr[:0:0], arr...), // Можно копией можно напрямую менять массив
	}
	num.sorted()
	fmt.Println(num.arr)
	//fmt.Println(arr)
}

func (q *QuickSort) sorted() {
	low := 0
	high := len(q.arr) - 1
	q.quickSort(low, high)
}

func (q *QuickSort) quickSort(low int, high int) {

	if low < high {
		// Выбираем опорный элемент
		pivotIndex := q.partition(low, high)

		// Рекурсивно сортируем подмассивы до и после опорного элемента
		q.quickSort(low, pivotIndex-1)
		q.quickSort(pivotIndex+1, high)
	}
}

func (q *QuickSort) partition(low, high int) int {
	// Используем последний элемент в качестве опорного
	pivot := q.arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		// Если текущий элемент меньше или равен опорному, меняем их местами
		if q.arr[j] <= pivot {
			i++
			q.arr[i], q.arr[j] = q.arr[j], q.arr[i]
		}
	}

	// Помещаем опорный элемент в его окончательную позицию
	q.arr[i+1], q.arr[high] = q.arr[high], q.arr[i+1]

	return i + 1
}
