package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type MyChannel struct {
	ch       chan string
	isClosed bool
	once     sync.Once
}

func (mc *MyChannel) SafeClose() {
	fmt.Println("SafeClose: сэйвово закрываю канал")
	mc.once.Do(func() {
		close(mc.ch)
		mc.isClosed = true
	})
}

func main() {
	var seconds int
	var workers int
	flag.IntVar(&seconds, "t", 10, "Количество секунд работы программы. дефолт: 10")
	flag.IntVar(&workers, "w", 5, "Количество воркеров программы. дефолт: 5")
	flag.Parse()
	fmt.Println("Программа завершится через", seconds, "секунд.")

	mc := &MyChannel{ch: make(chan string)}

	wg := &sync.WaitGroup{}
	wg.Add(4) // Инкриментим на 4 горутины
	for i := 0; i != workers; i++ {
		go sender(mc, wg)
		go reader(mc, wg)
	}

	// даём программе отработать указанное количество секунд
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("\n", "(main): время вышло")

	mc.SafeClose() // Сейвово закрываем канал

	wg.Wait() // Ожидаем выполнение всех горутин - пока wg.Add не приобритет значение 0
}

func sender(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("(sender): начинаю отправку данных в канал")
	for {
		if mc.isClosed {
			fmt.Println("(sender): канал закрыт, выходим")
			return
		}
		mc.ch <- "."
		time.Sleep(200 * time.Millisecond) // пауза 200 мс чтобы не спамить.
	}
}

func reader(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("(reader): начинаю читать канал")

	// range среагирует на закрытие канала и горутина завершится
	for data := range mc.ch {
		fmt.Print(data)
	}
	fmt.Println("(reader): читать больше нечего, выходим")
}
