// Переставить элементы в обратном порядке
package main

import "fmt"

func main() {
	arr := [10]string{}

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := len(arr) - 1; i > -1; i-- {
		fmt.Print(arr[i], " ")
	}
}
