package main

import "fmt"

func main() {
	var inputSlice = []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(findSet(inputSlice))

}

func findSet(inputSlice []string) []string {
	var outputSlice = []string{}

	m := make(map[string]int) // создаем мапу, куда добавляем элементы входного слайса в качестве ключа и их количество в качестве значения

	for _, val := range inputSlice {
		m[val] = 1
	}

	for val := range m { // находим уникальные элементы

		outputSlice = append(outputSlice, val)
	}
	return outputSlice
}
