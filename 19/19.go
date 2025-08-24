package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Println("Введите строку")

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = str[:len(str)-1] // удаляем символ новой строки

	fmt.Println(reverseStr(str))
}

func reverseStr(str string) string {

	strRune := []rune(str)
	l := len(strRune)
	for i := 0; i < l/2; i++ {
		strRune[i], strRune[l-i-1] = strRune[l-i-1], strRune[i]
	}
	str = string(strRune)

	return str
}
