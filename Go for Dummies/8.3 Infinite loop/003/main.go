// Счётчик чётных чисел
package main

import "fmt"

func main() {
	var (
		total  uint8
		number int
	)

	for {
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		if number&1 == 0 {
			total++
		} else {
			fmt.Println("Нечётное число!")
		}
	}

	fmt.Println(total)
}
