package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, c big.Int
	fmt.Println("Введите 2 числа")
	fmt.Scan(&a, &b)
	//a.Exp(big.NewInt(2), big.NewInt(80), nil)
	//b.Exp(big.NewInt(2), big.NewInt(40), nil)

	fmt.Printf("a + b = %v\n", c.Add(&a, &b))
	fmt.Printf("a - b = %v\n", c.Sub(&a, &b))
	fmt.Printf("a * b = %v\n", c.Mul(&a, &b))
	fmt.Printf("a / b = %v\n", c.Div(&a, &b))
}
