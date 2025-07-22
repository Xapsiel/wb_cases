package main

import "fmt"

type Summarizer interface { //целевой интерфейс, с которым работает клиент
	Summarize(numbers ...int) int
}

type Subtractor struct { //Адаптируемый объект, реализующий другой интерфейс
}

func (s *Subtractor) Subtract(numbers ...int) int {
	res := 0
	for i, n := range numbers {
		if i == 0 {
			res = n
			continue
		}
		res -= n
	}
	return res
}

type SubstractorAdapter struct { //Адаптер
	*Subtractor
}

func (s *SubstractorAdapter) Summarize(numbers ...int) int {
	numbers = append([]int{0}, numbers...)
	return s.Subtract(numbers...) * (-1) //адаптируем логику метода под интерфейс
}

func NewAdapter(adaptee *Subtractor) Summarizer {
	return &SubstractorAdapter{adaptee}
}

func main() {
	adapter := NewAdapter(&Subtractor{})
	res := adapter.Summarize(1, 2, 3)
	fmt.Println(res)
}
