// Количество элементов, равные наибольшему элементу
package main

import "fmt"

func main() {
	var (
		number int // Текущее число
		max    int // Максимум
		total  int // Число максимальных элементов
	)

	for {
		fmt.Scan(&number)
		if number == 0 {
			break
		}

		if number > max {
			max = number
			total = 1
		} else if number == max {
			total++
		}
	}

	fmt.Println(total)
}
