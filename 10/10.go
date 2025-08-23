package main

import "fmt"

func main() {

	var T = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var TMap = make(map[int][]float32)

	for _, value := range T {
		TMap[int(value/10)*10] = append(TMap[int(value/10)*10], value)
	}

	fmt.Println(TMap)
}
