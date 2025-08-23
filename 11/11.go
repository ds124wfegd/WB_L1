package main

import "fmt"

func main() {
	var firstSlice = []int{2, 3, 5, 4, 0}
	var secondSlice = []int{5, 4, 1, 0, 11, 22, 57, 0}

	fmt.Println(findIntersection(firstSlice, secondSlice))

}

func findIntersection(firstSlice, secondSlice []int) []int {
	var intersection = []int{}

	m := make(map[int]int) // создаем мапу, куда добавляем элементы первого слайса в качестве ключа и их количество в качестве значения

	for _, val := range firstSlice {
		m[val]++
	}

	for _, val := range secondSlice { // ищем пересечение
		if value, ok := m[val]; ok && (value > 0) {
			m[val]--
			intersection = append(intersection, val)
		}
	}
	return intersection
}
