// Сумма больше n
package main

import "fmt"

func main() {
	var n int
	var number int

	fmt.Scan(&n)

	sum := 0

	for sum <= n {
		fmt.Scan(&number)
		sum += number
	}

	fmt.Println(sum)
}
