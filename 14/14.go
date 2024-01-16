package main

import (
	"fmt"
	"reflect"
)

type testInterface interface {
	testMethod()
}

type testsStruct struct {
}

func (t testsStruct) testMethod() { fmt.Println("Test Method") }

func main() {
	ch := make(chan int)
	ts := testsStruct{}
	var st = testInterface(testsStruct{})
	arr := []any{1, "string", true, 1.2, ch, ts, st}
	for _, i := range arr {
		fmt.Println(
			func(variable interface{}) reflect.Type {
				return reflect.TypeOf(variable)
			}(i),
			"\n")
	}
}
