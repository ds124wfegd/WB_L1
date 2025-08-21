package main

import (
	"fmt"
	"sync"
)

func main() {
	// если не использовать wg, программа завершается раньше, чем произойдет выполнение горутин
	var wg sync.WaitGroup
	var arr = []int{2, 4, 6, 8, 10}

	for _, value := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(value * value)
		}()
	}
	wg.Wait()

}
