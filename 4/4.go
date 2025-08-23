package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, mainChan <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case num, ok := <-mainChan:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт, завершаю работу\n", id)
				return
			}
			fmt.Printf("worker %d обработал: %d\n", id, num)
			time.Sleep(300 * time.Millisecond)
		case <-ctx.Done():
			fmt.Printf("worker %d: получен сигнал завершения\n", id)
			return
		}
	}
}

func recording(mainChan chan<- int, ctx context.Context) {
	defer close(mainChan)

	for {
		// случайное целое число от 0 до 99
		i := rand.Intn(100)

		select {
		case <-ctx.Done():
			fmt.Println("producer: получен сигнал завершения, останавливаюсь")
			return
		default:

			select {
			//ожидание завершение контекста или отправки данных в канал
			case mainChan <- i:
				fmt.Println("отправлено число", i)
				time.Sleep(500 * time.Millisecond)
			case <-ctx.Done():
				fmt.Println("producer: получен сигнал завершения во время отправки")
				return
			}
		}
	}
}

func main() {
	var numWorker int
	fmt.Println("введите количество worker")
	fmt.Scanf("%d", &numWorker)

	if numWorker > 0 {
		fmt.Println("введите команду Ctrl+C для проверки корректного завершения")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// канал для обработки сигналов (graceful shutdown)
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			//сигнал завершения
			sig := <-sigChan
			fmt.Printf("получен сигнал %v\n", sig)
			cancel()
		}()

		mainChan := make(chan int, 10)
		var wg sync.WaitGroup

		for i := 1; i <= numWorker; i++ {
			wg.Add(1)
			go worker(i, mainChan, ctx, &wg)
		}

		recording(mainChan, ctx) // если поставить перед чтением из канала, получим deadlock

		wg.Wait()

		fmt.Println("выполнение окончено")

	} else {
		fmt.Println("Введите положительное число")
	}

}
