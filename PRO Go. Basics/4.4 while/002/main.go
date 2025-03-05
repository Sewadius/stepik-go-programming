// Список степеней двойки
package main

import "fmt"

func main() {
	const Two = 2
	var n int
	fmt.Scan(&n)

	number := 1

	for number <= n {
		fmt.Println(number)
		number *= Two
	}
}
