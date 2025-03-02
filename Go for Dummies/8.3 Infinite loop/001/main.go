// Бесконечная сумма
package main

import "fmt"

func main() {
	var (
		sum    int
		number int
	)

	for {
		fmt.Scan(&number)

		if number < 0 {
			break
		}

		sum += number
	}

	fmt.Println(sum)
}
