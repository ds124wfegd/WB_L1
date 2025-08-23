package main

import (
	"errors"
	"fmt"
)

/*
Основыные битовые операторы:
& -побитовое И
| - побитовое ИЛИ
a ^ b  - побитовое исключающее ИЛИ
^a - побитовое НЕ
<< - сдвиг влево
>> - сдвиг вправо
*/

func main() {
	var num, res, i int64
	var err error
	fmt.Println("Введите, число и номер бита, который необходимо изменить")
	fmt.Scanf("%d %d", &num, &i)

	res, err = changeIBit(num, i-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("При изменении %d бита получим %d\n", i, res)
	}
}

func changeIBit(num, i int64) (res int64, err error) {
	if num < (1 << i) {
		return 0, errors.New("y числа отсутствует указанный бит")
	} else {
		return num ^ (1 << i), nil
	}

}
