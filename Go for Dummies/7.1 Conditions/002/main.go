// Арифметическая прогрессия
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c-b == b-a {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
