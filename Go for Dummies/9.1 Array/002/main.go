// Найти 2 максимума в массиве
package main

import "fmt"

func main() {
	var (
		arr    [20]int // Массив чисел
		index1 int     // Позиция первого максимума
		index2 int     // Позиция второго максимума
	)

	// Чтение чисел
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	// Начальные значения для максимумов
	max1, max2 := arr[0], arr[1]

	// Сохранение первого максимума и его позиции
	for i := 0; i < len(arr); i++ {
		if arr[i] > max1 {
			max1, index1 = arr[i], i
		}
	}

	// Сохранение второго максимума и его позиции
	for i := 0; i < len(arr); i++ {
		if arr[i] > max2 && i != index1 {
			max2, index2 = arr[i], i
		}
	}

	// Вывод в порядке появления в массиве
	if index1 < index2 {
		fmt.Println(max1, max2)
	} else {
		fmt.Println(max2, max1)
	}
}
