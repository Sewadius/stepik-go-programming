// Калькулятор
package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		Add  = "+"
		Sub  = "-"
		Mult = "*"
		Div  = "/"
		Rem  = "%"
	)

	var (
		a, b, result float32
		oper         string
	)

	fmt.Scan(&a, &b, &oper)

	if oper == Div && b == 0 {
		fmt.Println("Делить на ноль нельзя!")
		os.Exit(0)
	}

	switch oper {
	case Add:
		result = a + b
	case Sub:
		result = a - b
	case Mult:
		result = a * b
	case Div:
		result = a / b
	}

	if oper == Rem {
		fmt.Println(int(a) % (int(b)))
	} else {
		fmt.Println(result)
	}
}
