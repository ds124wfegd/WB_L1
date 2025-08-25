package main

import "fmt"

/* используем дженерики, это имеет следующие преимущества:

Повторное использование кода - одна функция работает с любыми типами

Type safety - компилятор проверяет типы во время компиляции

Избегание interface{} и type assertions - меньше runtime ошибок
*/
func removeElement[T any](slice []T, index int) []T {
	// Проверка корректности индекса
	if index < 0 || index >= len(slice) {
		println("Индекс выходит за пределы массива")
		return slice
	}

	// Сдвигаем элементы на место удаляемого
	copy(slice[index:], slice[index+1:])

	// Обнуляем последний элемент (для предотвращения утечки памяти)
	var zero T
	slice[len(slice)-1] = zero

	// Возвращаем срез с уменьшенной длиной
	return slice[:len(slice)-1]
}

func main() {
	// Исходный слайс
	var i int
	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println("Укажите, индекс символа в массиве, который необходимо удалить")
	fmt.Scanf("%d\n", &i)

	fmt.Println("Original slice:", numbers)

	// Удаляем элемент с индексом i
	numbers = removeElement(numbers, i)
	fmt.Printf("После удаления элемента c индексом %d: %v\n", i, numbers)

}
