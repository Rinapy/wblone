package main

import "fmt"

type Human struct {
	Name string
	Age  int
	Sex  string
}

type Action struct {
	*Human
}

func (a *Action) sayName() {
	fmt.Printf("Hello, my name %s.", a.Name)
}

func main() {
	ac := Action{
		&Human{
			Name: "Bill",
			Age:  23,
			Sex:  "male",
		},
	}
	ac.sayName()
}
