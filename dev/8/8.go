package main

import "fmt"

func main() {
	var num int64 = 25  // Исходное число
	position := uint(3) // Позиция бита, который хотим изменить
	bitVal := int64(0)  // Желаемое значение бита (0 или 1)

	result := setBit(num, position, bitVal)
	fmt.Printf("Исходное значение: %d\n", num)
	fmt.Printf("Измененное значение: %d\n", result)
}

func setBit(num int64, i uint, bitVal int64) int64 {
	mask := int64(1 << i) // Создаем маску с помощью побитового сдвига
	num &= ^mask          // Обнуляем i-й бит в числе, применяя побитовое отрицание на маску
	num |= bitVal << i    // Устанавливаем i-й бит в заданное значение
	return num
}
