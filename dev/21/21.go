package main

import (
	"fmt"
)

// Интерфейс ожидаемого клиентом класса
type Target interface {
	Request() string
}

// Адаптируемый класс
type Adaptee struct{}

// Реквест который возвращает rune
func (a *Adaptee) Request() []rune {
	str := "Специфический запрос"
	r := []rune(str)
	return r
}

// Адаптер
type Adapter struct {
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	specificRequest := a.Adaptee.Request()
	// Выполняем необходимые преобразования и логику для адаптации запроса
	return fmt.Sprintf("Адаптированный запрос: %s", string(specificRequest))
}

type Adapter2 struct {
	Adaptee *Adaptee
}

func (a *Adapter2) Request() string {
	specificRequest := a.Adaptee.Request()
	upd := string(specificRequest[0:5]) + string(specificRequest[4])
	// Выполняем необходимые преобразования и логику для адаптации запроса
	return fmt.Sprintf("Адаптированный запрос: %s", string(upd))
}

func main() {
	adaptee := &Adaptee{}
	adapter1 := &Adapter{Adaptee: adaptee}
	adapter2 := &Adapter2{Adaptee: adaptee}
	fmt.Println(adaptee.Request())
	target := Target(adapter1)
	fmt.Println(target.Request())
	target = Target(adapter2)
	fmt.Println(target.Request())
	// Используем адаптер для выполнения запроса через интерфейс Target
}
