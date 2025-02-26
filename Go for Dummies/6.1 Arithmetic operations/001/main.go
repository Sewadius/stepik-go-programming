package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	fmt.Println("Частное:", a/b)
}
