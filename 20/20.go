package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseWords(s string) string {

	words := strings.Split(s, " ")
	l := len(words)

	for i := 0; i < l/2; i++ {
		words[i], words[l-i-1] = words[l-i-1], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("ошибка при чтении строки", err)
	}
	s = strings.TrimSpace(s) // удаляем пробельные символы с начала и конца строки
	fmt.Println(reverseWords(s))
}
