/*
var justString string

func someFunc() {
	v := createHugeString(1 << 10)  //создается строка размером 1024 байта
	justString = v[:100]  // из строки берется 100 первых символов, сохраняется ссылка на срез строки v (базовый массив - строка v)
						// поэтому при завершении someFunc() значение переменной v не будет очищено Чтобы решить эту проблему, нужно
						// создать копию подстроки, чтобы она больше не ссылалась на оригинальную строку
}

func main() {
	someFunc()
}

*/

package main

import "fmt"

func createString(size int) string {
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createString(1 << 10)
	justString = string(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
