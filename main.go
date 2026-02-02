package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	jobs := make(chan int)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		for i := 1; ; i++ {
			time.Sleep(3 * time.Second)
			jobs <- i
		}
	}()
	for {
		select {
		case job := <-jobs:
			fmt.Printf("Обработка задачи №%d\n", job)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Готово!")
		case s := <-sigs:
			fmt.Printf("\n Получен сигнал: %v\n", s)
			time.Sleep(2 * time.Second)
			fmt.Println("Все данные сохранены. Выход.")
			return

		case <-time.After(2 * time.Second):
			// Это сработает, если в jobs ничего не приходило 2 секунды
			fmt.Println("логов нет, ждем новых задач")
		}
	}
}
