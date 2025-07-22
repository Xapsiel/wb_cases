package main

import (
	"fmt"
	"math/big"
)

func Action(a, b *big.Int) {
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s + %s = %s\n", a, b, sum)

	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Разность: %s - %s = %s\n", a, b, diff)

	prod := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s * %s = %s\n", a, b, prod)

	if b.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Деление: деление на ноль")
	} else {
		quot := new(big.Int).Div(a, b)
		fmt.Printf("Деление: %s / %s = %s\n", a, b, quot)
	}
}

func main() {
	a := big.NewInt(1 << 62)
	b := big.NewInt(1 << 42)
	fmt.Printf("Числа: a = %s, b = %s\n", a, b)
	performBigOperations(a, b)
}
