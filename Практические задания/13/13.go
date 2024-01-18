package main

import "fmt"

func main() {
	a, b := "first", "second"
	c, d := 1, 2
	fmt.Println("Было ", a, b, " - ", c, d)

	swap(&a, &b)
	swap(&c, &d)

	fmt.Println("Стало ", a, b, " - ", c, d)

}

func swap[T any](first, second *T) {
	*first, *second = *second, *first
}
