package main

import "fmt"

func main() {
	var a, b float32

	fmt.Scan(&a, &b)

	fmt.Println("Сложение:", a+b)
	fmt.Println("Вычитание:", a-b)
	fmt.Println("Умножение:", a*b)
	fmt.Println("Деление:", a/b)
}
