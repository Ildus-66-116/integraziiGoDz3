package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Задание 2. Graceful shutdown
// -В работе часто возникает потребность правильно останавливать приложения.
// Например, когда наш сервер обслуживает соединения, а нам хочется,
// чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса.
// Для этого существует паттерн graceful shutdown.
// -Напишите приложение, которое выводит квадраты натуральных чисел на экран,
// а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.
func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("%d^2 = %d\n", i, i*i)
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-stop:
		fmt.Println("выхожу из программы")
	}
}
