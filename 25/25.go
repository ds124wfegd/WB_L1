package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func Sleep2(duration time.Duration) {
	done := make(chan struct{})

	go func() {
		// имитируем ожидание без time.Sleep
		start := time.Now()
		for time.Since(start) < duration {
			// пустой цикл ожидания
		}
		close(done)
	}()

	<-done
}

func main() {
	fmt.Println("Использование канала c time.After")
	start := time.Now()
	sleep(2 * time.Second)
	fmt.Println("c момента старта прошло", time.Since(start))

	fmt.Println("Использование горутины и канала")
	start = time.Now()
	Sleep2(3 * time.Second)
	fmt.Println("c момента старта прошло", time.Since(start))
}
