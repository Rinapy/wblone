package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := Counter{
		i: 0,
	}

	var wg sync.WaitGroup
	// 10 на инкримент
	for i := 0; i != 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	// 3 на декремент
	for i := 0; i != 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Dinc()
		}()
	}
	wg.Wait()
	fmt.Println(counter.i) // 7
}

type Counter struct {
	i  int
	mu sync.Mutex // для устранения гонки данных
}

func (c *Counter) Inc() {
	c.mu.Lock() // лочим
	c.i++
	c.mu.Unlock() // унлочим
}

func (c *Counter) Dinc() {
	c.mu.Lock() // лочим
	c.i--
	c.mu.Unlock() // унлочим
}
