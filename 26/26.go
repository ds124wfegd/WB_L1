package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	m := make(map[rune]struct{}) //struct{} (пустая структура - занимает 0 байт памяти)
	for _, simbol := range s {
		if _, ok := m[simbol]; ok {
			return false
		}
		m[simbol] = struct{}{}
	}
	return true
}

func main() {
	var str string
	fmt.Println("Введите строку")
	fmt.Scanf("%s\n", &str)
	fmt.Println(isUnique(strings.ToLower(str)))

}
