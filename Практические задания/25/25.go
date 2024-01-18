package main

import (
	"fmt"
	"time"
)

func CustomSleep(duration time.Duration) {
	<-time.After(duration) // принимаем значение с помощью оператора <- и блокируем выполнение горутины на указанное время.
}

func main() {
	fmt.Println("Вызываем слип")
	CustomSleep(12 * time.Second)
	fmt.Println("Печатаем после слипа")
}
