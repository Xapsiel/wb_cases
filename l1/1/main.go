package main

import "fmt"

type Human struct {
	Name    string
	SubName string
	Age     int
	Sex     string
}

type Action struct {
	*Human
	Action string
}

func NewAction(human *Human) *Action {
	return &Action{
		Human: human,
	}
}

func NewHuman(name, subname, sex string, age int) *Human {
	return &Human{
		Name:    name,
		SubName: subname,
		Age:     age,
		Sex:     sex,
	}
}

func (h *Human) BirthDay() {
	h.Age++
}

func (a *Action) Set(action string) {
	a.Action = action
}
func (a *Action) Do() {
	fmt.Printf("%s %s выполняет следующее: %s \n", a.Name, a.SubName, a.Action)
}
func main() {
	human := NewHuman("Иван", "Иванов", "Мужчина", 17)
	a := NewAction(human)
	a.BirthDay()
	a.Set(fmt.Sprintf("Празднует свой %d день рождение", human.Age))
	a.Do()

}
