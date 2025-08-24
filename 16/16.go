package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr = []int{5, 0, 3, -1, 6, 8, 55, 2, 1, 444, 33, 0, 1, -8, -8}

	fmt.Println(quickSort(arr))

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
