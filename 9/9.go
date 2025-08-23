package main

import "fmt"

func main() {
	numbers := []int{12, 89, 56, 632, 0, 32, 56, 42, 63, 96}

	ch1 := make(chan int) // канал для записи исходных чисел из массива
	ch2 := make(chan int) // канал для записи результата операции x*2

	go func() { // запись чисел в канал
		defer close(ch1)
		for _, x := range numbers {
			ch1 <- x
		}
	}()

	go func() { // чтение из канала 1, умножение на 2, запись во 2-ой канал
		defer close(ch2)
		for x := range ch1 {
			ch2 <- x * 2
		}
	}()

	// Чтение из второго канала и вывод в stdout
	for result := range ch2 {
		fmt.Println(result)
	}
}
