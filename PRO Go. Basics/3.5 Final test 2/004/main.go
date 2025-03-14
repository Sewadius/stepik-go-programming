package main

import "fmt"

func main() {
	var n float32
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Обратного числа не существует")
	} else {
		fmt.Println(1 / n)
	}
}
