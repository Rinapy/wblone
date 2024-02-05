package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Введите набор слов для их разворота:")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// чистим ненужные символы перевода каретки
	str := strings.ReplaceAll(input, "\r\n", "")
	fmt.Printf("%q \n", str)
	sliceOfStrings := strings.Split(str, " ")
	slices.Reverse(sliceOfStrings)

	fmt.Println(strings.Join(sliceOfStrings, " "))
}
