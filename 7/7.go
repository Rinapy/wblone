package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	ma map[int]string
	mu sync.RWMutex
}

func (m *MyMap) WriteToMap(key int, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ma[key] = value
}

func main() {
	mymap := MyMap{
		ma: make(map[int]string),
		mu: sync.RWMutex{},
	}
	var wg sync.WaitGroup
	for i := 0; i != 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mymap.WriteToMap(i, "запись в мапе")
		}(i)
	}
	wg.Wait()

	for key, value := range mymap.ma {
		fmt.Printf("Горутина - %v: %s\n", key, value)
	}
	fmt.Println("Длинна мапы: ", len(mymap.ma))

}
