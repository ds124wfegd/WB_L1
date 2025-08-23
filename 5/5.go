package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var timeWorking int64
	fmt.Println("Введите время работы программы (в секундах)")
	fmt.Scanf("%d\n", &timeWorking)

	// Создаем канал
	channel := make(chan int)

	timer := time.NewTimer(time.Duration(timeWorking) * time.Second)
	defer timer.Stop()

	// Горутина для отправки данных
	go func() {

		for {
			num := rand.Intn(100)
			select {
			case channel <- num:
				fmt.Printf("отправлено число %d\n", num)
				time.Sleep(500 * time.Millisecond)
			case <-timer.C:
				fmt.Println("таймер истек, отправка закончена!")
				close(channel)
				return
			}
		}
	}()

	// Чтение данных
	for {
		data, ok := <-channel
		if !ok {
			fmt.Println("канал закрыт")
			return
		}
		fmt.Printf("принято число: %d\n", data)
	}
}
