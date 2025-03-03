// Поиск числа Фибоначчи
package main

import "fmt"

func main() {
	var number int
	a, b := 1, 1

	fmt.Scan(&number)

	for a <= number {
		b, a = a+b, b
	}

	fmt.Println(a)
}
