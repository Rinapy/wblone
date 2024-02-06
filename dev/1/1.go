package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Sex   string
	Power int
}

type Action struct {
	*Human
}

func (a *Action) sayName() {
	fmt.Printf("Hello, my name %s.\n", a.Name)
}

func (a *Action) goTraining() {
	fmt.Printf("%s, go to training.\n", a.Name)
	a.Power--
}

func main() {
	hum := Human{
		Name: "Bill",
		Age:  23,
		Sex:  "male",
	}
	ac := Action{
		&hum,
	}
	ac.sayName()
	ac.goTraining()
}
