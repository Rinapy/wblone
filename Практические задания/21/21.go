package main

import "fmt"

// Интерфейс ожидаемого клиентом класса
type Target interface {
	Request() string
}

// Адаптируемый класс
type Adaptee struct{}

// Реквест который возвращает rune
func (a *Adaptee) SpecificRequest() []rune {
	str := "Специфический запрос"
	r := []rune(str)
	return r
}

// Адаптер
type Adapter struct {
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	specificRequest := a.Adaptee.SpecificRequest()
	// Выполняем необходимые преобразования и логику для адаптации запроса
	return fmt.Sprintf("Адаптированный запрос: %s", string(specificRequest))
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{Adaptee: adaptee}

	// Используем адаптер для выполнения запроса через интерфейс Target
	result := adapter.Request()
	fmt.Println(result)
}
