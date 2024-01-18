package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	supStr := "Запуск и остановка через "
	// Chan
	// Плюсы:
	// - Каналы предоставляют механизм синхронизации и обмена данными между горутинами.
	// - Блокирующие операции чтения (receive) из канала могут вызывать приостановку горутины, пока данные не станут доступны.
	// - Закрытие канала может использоваться для уведомления горутин о необходимости остановки.
	// Минусы:
	// - Неправильное использование каналов может привести к блокировке или утечке ресурсов.
	// - Закрытие канала не немедленно завершает все чтения и записи, поэтому необходимо корректно управлять остановкой горутин.
	fmt.Println(supStr, Chan(), "\n")
	time.Sleep(7 * time.Second)

	// Context
	// Плюсы:
	//- Контекст предоставляет средства для управления временем жизни и прерывания работы горутин.
	//- Отмена контекста автоматически распространяется на все производные контексты и связанные с ними операции.
	//- Позволяет передавать дополнительные значения и состояния через контекст.
	//
	// Минусы:
	//- Пропагация контекста может стать громоздкой и усложнить код.
	//- Неправильное использование контекста может привести к неожиданному поведению и утечкам ресурсов.
	fmt.Println(supStr, Context(), "\n")
	time.Sleep(7 * time.Second)

	// Flag
	// Плюсы:
	//- Простой и непосредственный способ контролировать выполнение горутин.
	//- Позволяет осуществлять мягкую остановку (graceful shutdown) путем изменения значений переменных состояния.
	//
	// Минусы:
	//- Может быть сложно гарантировать правильную синхронизацию при доступе к разделяемым переменным.
	//- Не предоставляет механизмы автоматической блокировки или уведомления горутин о необходимости остановки.
	fmt.Println(supStr, Flag())
	time.Sleep(7 * time.Second)

	// Timer
	fmt.Println(supStr, Timer(), "\n")
	time.Sleep(7 * time.Second)

	//Time
	fmt.Println(supStr, Time(), "\n")
	time.Sleep(11 * time.Second)

	// OS Signals
	// Плюсы:
	//- Сигналы предоставляют механизм взаимодействия с операционной системой и могут использоваться для остановки горутин.
	//- Позволяют гибко и явно управлять процессом завершения горутины.
	//
	// Минусы:
	//- Необходимо быть осторожным при обработке сигналов, чтобы избежать гонок данных и других проблем синхронизации.
	//- Не все операционные системы имеют одинаковую поддержку функций сигналов.
	fmt.Println(supStr, OSSignal())

}

func Chan() string {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Остановка горутины через канал...")
				return
			default:
				fmt.Println("Канальная горутина работает...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	return "Chan"
}

func Context() string {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Остановка горутины через context...")
				return
			default:
				fmt.Println("Контекстаня горутина работает...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(4 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	return "Context"
}

func Flag() string {
	var flag bool = false
	go func() {
		for {
			if flag {
				fmt.Println("Остановка горутины через flag...\n")
				return
			}
			fmt.Println("Флаг горутина работает...")
			time.Sleep(time.Second)

		}
	}()
	time.Sleep(6 * time.Second)
	flag = true
	return "Flag"
}

func OSSignal() string {
	// Создаем канал для получения сигналов
	signals := make(chan os.Signal, 1)
	// Регистрируем канал для получения сигналов SIGINT и SIGTERM
	signal.Notify(signals, os.Interrupt, syscall.SIGINT)
	// Запускаем горутину
	go func() {
		for {
			fmt.Println("Сигнал горутина работает...")
			time.Sleep(1 * time.Second)
		}
	}()
	<-signals
	time.Sleep(3 * time.Second)
	//Имитируем сигнал
	process, _ := os.FindProcess(os.Getpid())
	process.Signal(os.Interrupt) // Не достаточно корректно работает на винде или я что-то не понял

	return "OSSignal"
}

func Timer() string {
	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fmt.Println("Остановка горутины через timer...")
				return
			default:
				fmt.Println("Таймер горутина работает...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(6 * time.Second)
	return "Timer"
}

func Time() string {
	go func() {
		fmt.Println("Остановка горутины через time...")
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		fmt.Println("Тайм горутина отработала...")
	}()
	time.Sleep(11 * time.Second)
	return "Time"
}
