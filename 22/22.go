package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Словил на винде Access is denied
	a := big.NewInt(12254412432322)
	b := big.NewInt(512612313)

	// Сложение
	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Println("Сложение:", sum)

	// Вычитание
	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Println("Вычитание:", sub)

	// Умножение
	mul := big.NewInt(0)
	mul.Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := big.NewInt(0)
	if b.Cmp(big.NewInt(0)) != 0 { // ограничение деления на 0
		div.Div(a, b)
	}
	fmt.Println("Деление:", div)

}
