// Сумма
package main

import "fmt"

func main() {
	var n, sum, number int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		sum += number
	}

	fmt.Println(sum)
}
