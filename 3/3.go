package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var numWorker int
	fmt.Println("Введите количество worker")
	fmt.Scanf("%d", &numWorker)

	if numWorker > 0 {

		mainChan := make(chan int, 10)
		var wg sync.WaitGroup

		for i := 1; i <= numWorker; i++ {
			wg.Add(1)
			go worker(i, mainChan, &wg)
		}

		recording(mainChan) // если поставить перед чтением из канала, получим deadlock

		wg.Wait()

		fmt.Println("выполнение окончено")

	} else {
		fmt.Println("Введите положительное число")
	}

}

func recording(mainChan chan<- int) {
	defer close(mainChan)

	for {
		// Случайное целое число от 0 до 99
		i := rand.Intn(100)
		mainChan <- i
		// Пауза между отправками
		time.Sleep(500 * time.Millisecond)
		fmt.Println("отправлено число", i)
	}
}

func worker(id int, mainChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range mainChan {

		fmt.Printf("worker %d обработал: %d\n", id, num)
	}
}
