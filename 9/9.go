package main

import (
	"fmt"
	"time"
)

func main() {
	numsCh := make(chan int)
	resCh := make(chan int)
	nums := make([]int, 0)
	for i := 0; i != 20; i++ {
		nums = append(nums, i+1)
	}

	go ReadToSqr(numsCh, resCh)
	go SqrToStd(resCh)
	for _, v := range nums {
		numsCh <- v
	}
}

func ReadToSqr(numsCh chan int, resCh chan int) {
	for {
		select {
		case x := <-numsCh:
			resCh <- x * x
		default:
			fmt.Println("в numsCh не осталось данных")
			time.Sleep(time.Second)
		}
	}
}

func SqrToStd(resCh chan int) {
	for {
		select {
		case sqr := <-resCh:
			fmt.Println(sqr, "\t")
			time.Sleep(time.Second) // Чтобы не заспамить
		default:
			fmt.Println("в resCh не осталось данных")
			time.Sleep(time.Second)
		}
	}

}
