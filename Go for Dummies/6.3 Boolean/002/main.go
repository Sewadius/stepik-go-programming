// Проверка на школьника
package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)

	result := age > 12 && age < 20
	fmt.Println(result)
}
