package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func workerWithCondition(stopFlag *bool) {
	for !*stopFlag {
		fmt.Println("Работаю...")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Горутина завершена по условию")
}

func workerWithChannel(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Горутина завершена по сигналу канала")
			return
		default:
			fmt.Println("Выполняю работу...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина завершена по контексту:", ctx.Err())
			return
		default:
			fmt.Println("Работаю с контекстом...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func workerWithGoexit() {
	defer fmt.Println("Горутина завершена через Goexit()")

	for {
		fmt.Println("Работаю... (остановлюсь через 2 секунды)")
		time.Sleep(500 * time.Millisecond)

		// Эмуляция условия для остановки
		if time.Now().Unix()%5 == 0 { // Условие для демонстрации
			runtime.Goexit()
		}
	}
}

func workerWithPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Горутина восстановлена после паники:", r)
		}
	}()

	for {
		fmt.Println("Работаю... (упаду через 2 секунды)")
		time.Sleep(500 * time.Millisecond)

		if time.Now().Unix()%4 == 0 { // Условие для демонстрации
			panic("экстренная остановка!")
		}
	}
}

func workerWithWG(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Printf("Воркер %d: итерация %d\n", id, i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}

func workerWithDataChannel(dataChan chan int) {
	for data := range dataChan { // Цикл завершится при закрытии канала
		fmt.Printf("Обрабатываю данные: %d\n", data)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Горутина завершена (канал данных закрыт)")
}

// worker обрабатывает данные из канала с возможностью остановки и таймаута
func worker(stopChan <-chan struct{}, dataChan <-chan int, timeoutChan <-chan time.Time) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Горутина: получила сигнал остановки")
			return
		case <-timeoutChan:
			fmt.Println("Горутина: таймаут истек")
			return
		case data, ok := <-dataChan:
			if !ok {
				fmt.Println("Горутина: канал данных закрыт")
				return
			}
			fmt.Printf("Горутина: получила данные %d\n", data)
		default:
			fmt.Println("Горутина: работает...")
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func main() {

	// ------------выход по условию------------
	stop := false
	go workerWithCondition(&stop)
	time.Sleep(2 * time.Second)
	stop = true // Устанавливаем флаг остановки
	time.Sleep(100 * time.Millisecond)

	// ------------выход через канал------------
	stopChan := make(chan struct{})
	go workerWithChannel(stopChan)
	time.Sleep(2 * time.Second)
	close(stopChan) // Закрытие канала - сигнал к остановке
	time.Sleep(100 * time.Millisecond)

	// ------------через контекст------------

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go workerWithContext(ctx)

	time.Sleep(3 * time.Second) // Ждем завершения по таймауту

	// ------------через runtime.Goexit() - аварийная ------------
	go workerWithGoexit()
	time.Sleep(3 * time.Second)

	// ------------через панику с восстановлением------------
	go workerWithPanic()
	time.Sleep(3 * time.Second)

	// ------------через waitGropup------------

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Увеличиваем счетчик
		go workerWithWG(&wg, i)
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("Все горутины завершены")

	// ------------через закрытие канала------------
	dataChan := make(chan int, 10)
	go workerWithDataChannel(dataChan)

	for i := 1; i <= 5; i++ { // Отправляем данные
		dataChan <- i
	}

	time.Sleep(1 * time.Second)
	close(dataChan) // Закрываем канал - сигнал к остановке
	time.Sleep(100 * time.Millisecond)

	// ------------с несколькими каналами------------

	stopChan2 := make(chan struct{})
	dataChan2 := make(chan int, 5)
	timeoutChan := time.After(3 * time.Second)

	go worker(stopChan2, dataChan2, timeoutChan) // Запускаем горутину с worker

	for i := 1; i <= 3; i++ { // Отправляем данные
		dataChan2 <- i
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
	close(stopChan2) // Останавливаем горутину
	time.Sleep(500 * time.Millisecond)

}
