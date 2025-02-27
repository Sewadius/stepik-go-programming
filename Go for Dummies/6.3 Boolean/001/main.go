// Четные или нечетные числа
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result := a%2 == 0 && b%2 == 0
	fmt.Println(result)
}
