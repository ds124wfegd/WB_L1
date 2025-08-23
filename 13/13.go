package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = XORMethod(a, b)
	fmt.Printf("a = %d, b = %d - после XORMethod\n", a, b)
	a, b = addMethod(a, b)
	fmt.Printf("a = %d, b = %d - после addMethod\n", a, b)
}

func XORMethod(a, b int) (int, int) { //метод сложения

	a = a ^ b // a = 15 (5 XOR 10)
	b = a ^ b // b = 5 (15 XOR 10)
	a = a ^ b // a = 10 (15 XOR 5)

	return a, b
}

func addMethod(a, b int) (int, int) { //метод сложения

	a = a + b // a = 15 (5 + 10)
	b = a - b // b = 5 (15 - 10)
	a = a - b // a = 10 (15 - 5)

	return a, b
}
