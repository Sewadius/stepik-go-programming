// Количество чисел
package main

import "fmt"

func main() {
	var n, number, total int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number%10 == 0 {
			total++
		}
	}

	fmt.Println(total)
}
