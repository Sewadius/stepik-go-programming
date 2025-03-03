// Чётное, нечётное, чётное, нечётное
package main

import "fmt"

func main() {
	const size = 15
	var (
		numbers             [size]int // Массив чисел
		even                [size]int // Чётные числа
		odd                 [size]int // Нечётные числа
		evenIndex, oddIndex int       // Индексы для массивов
	)

	// Чтение данных
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	// Распределение чисел по массивам
	for i := 0; i < size; i++ {
		if numbers[i]%2 == 0 {
			even[evenIndex] = numbers[i]
			evenIndex++
		} else {
			odd[oddIndex] = numbers[i]
			oddIndex++
		}
	}

	// Индексы для чётных и нечётных чисел
	i, j := 0, 0

	// Чередуем чётные и нечётные числа
	for i < evenIndex && j < oddIndex {
		fmt.Print(even[i], " ")
		i++
		fmt.Print(odd[j], " ")
		j++
	}

	// Добавляем оставшиеся чётные числа, если есть
	for ; i < evenIndex; i++ {
		fmt.Print(even[i], " ")
	}

	// Добавляем оставшиеся нечётные числа, если есть
	for ; j < oddIndex; j++ {
		fmt.Print(odd[j], " ")
	}

	fmt.Println()
}
