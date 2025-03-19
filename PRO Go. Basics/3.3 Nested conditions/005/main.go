// Калькулятор
package main

import "fmt"

func main() {
	var a, b int
	var operation string

	fmt.Scan(&a, &b, &operation)

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(float64(a) / float64(b))
		}
	default:
		fmt.Println("Неверная операция")
	}
}
