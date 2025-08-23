package main

import (
	"fmt"
	"reflect"
)

func defineType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}

func main() {
	fmt.Println(defineType(1))
	fmt.Println(defineType("str"))
	fmt.Println(defineType(true))
	fmt.Println(defineType(make(chan int)))
}
