// Таблица умножения
package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	for i := 1; i < 11; i++ {
		fmt.Println(number, "x", i, "=", i*number)
	}
}
