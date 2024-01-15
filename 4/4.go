package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// В данном примере myChan является каналом, который имеет ёмкость 1 (буферизованный канал).
// Ёмкость указывает, сколько элементов канала может содержать без блокировки отправителя.
// Однако в данном случае мы используем пустую структуру struct{}, так как нам не нужно передавать какие-либо конкретные значения через этот канал.
// Вместо этого, мы используем канал myChan в качестве сигнала для синхронизации, чтобы уведомить другую горутину о наступлении определенного события.
// Использование пустой структуры struct{} позволяет нам создать канал без каких-либо конкретных данных и минимизировать потребление памяти для него.
type myChan chan struct{}

func main() {

	ch := make(myChan, 1)

	var num int
	// парсинг количества воркеров из аргумента командной строки -w
	flag.IntVar(&num, "w", 10, "Введите число - количество воркеров.")
	flag.Parse()

	// переменная context получит уведомление от операционной системы
	// о нажатии ctrl+c пользователем. Этот сигнал далее передаётся всем,
	// кто подписан на данный контекст.
	ctx, _ := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	fmt.Println("Эта программа завершится по нажатию клавиш Ctrl+c")
	//Создание указанного пользователем количества воркеров
	//Для каждого воркера запускается функция reader(c) в отдельной горутине
	for i := 0; i < num; i++ {
		go reader(ch, i)
	}

	//запускает в новой горутине функцию writer, которая
	//будет записывать данные в канал
	writer(ch, ctx)
}

func writer(ch myChan, ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //  Когда канал будет закрыт, что указывает на отмену контекста, оператор <- будет разблокирован и выполнится ветка case <-ctx.Done().
			fmt.Println("\n writer: Graceful Shutdown по нажатию Ctrl+c")
			// закрытие канала остановит циклы for-range во всех reader'ах,
			// благодаря чему все читающие горутины сэйфово завершатся без утечек памяти.
			close(ch)
			return
		default:
			time.Sleep(100 * time.Millisecond) // пауза чтобы не заспамить консоль
			ch <- struct{}{}                   // отправляем в канал пустоту
		}
	}
}

func reader(ch myChan, id int) {
	for range ch {
		fmt.Print(id, " ") //
	}
}
