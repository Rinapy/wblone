package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// чистим ненужные символы перевода каретки
	str := strings.ReplaceAll(input, "\r\n", "")
	fmt.Printf("%q \n", str)
	runes := []rune(str)
	fmt.Printf("%q \n", runes)
	slices.Reverse(runes)

	fmt.Println(string(runes))
}
