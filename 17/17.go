package main

import (
	"fmt"
	"math/rand"
)

// итеративная реализация бинарного поиска
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // возвращаем найденный индекс
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // элемент не найден
}

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivotIndex := rand.Intn(len(arr))
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	var arr = []int{5, 0, 3, -1, 6, 8, 55, 2, 1, 444, 33, 0, 1, -8, -8}
	var num int
	fmt.Println("ведите число, которое необходимо найти")
	fmt.Scanf("%d\n", &num)
	quickSort(arr)
	fmt.Println("данное число имеет индекс", binarySearch(arr, num))

}
