package main

import "fmt"

//структура human - объявлением
type Human struct {
	Name    string
	Surname string
	Age     int
}

// методы структуры Human
func (h *Human) Run(speed float64) string {
	return fmt.Sprintf("%s, бегает co скоростью %f", h.Name, speed)
}

func (h *Human) Jump(height float64) string {
	return fmt.Sprintf("%s, прыгает на высоту %f", h.Name, height)
}

//структура Action - объявление
type Action struct {
	ActionDesc string
	Human
}

// метод структуры Action
func (a *Action) AddDesc(desc string) string {
	a.ActionDesc = desc
	return fmt.Sprintf("добавлено описание %s", a.ActionDesc)
}

func (a *Action) Jump(height, distance float64) string {
	return fmt.Sprintf("%s, прыгает на высоту %f и расстояние %f", a.Human.Name, height, distance)
}

func main() {

	act := new(Action)
	act.Human.Name = "Vasya"
	act.Human.Surname = "Ivanov"
	act.Human.Age = 22

	// допустимы вызовы методов
	fmt.Println(act.AddDesc("описание Action"))
	fmt.Println(act.Human.Run(20.1))
	fmt.Println(act.Run(14.2))
	//act.Human.Jump(20.1) - аналогично Jump
	//act.Jump(14.2) - аналогично Jump, если нет методы Jump у Action

	// но родительский тип может переопределить метод, причем метод перекрывается по имени, без учета полной сигнатуры метода:
	//fmt.Println(act.Jump(78.2))  // - некорректно
	fmt.Println(act.Jump(78.2, 55))  // - корректно
	fmt.Println(act.Human.Jump(4.2)) // - корректно

}
