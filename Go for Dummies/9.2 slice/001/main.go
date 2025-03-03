// Учитель
package main

import "fmt"

func main() {
	const FIVE = 5

	var (
		n               int // Количество оценок
		sum             int // Общая сумма оценок
		mark            int // Текущая оценка
		total_excellent int // Количество отличных оценок
	)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&mark)

		if mark == FIVE {
			total_excellent++
		}

		sum += mark
	}

	fmt.Println(total_excellent)
	fmt.Println(float32(sum) / float32(n))
}
