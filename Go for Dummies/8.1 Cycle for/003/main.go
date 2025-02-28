// Сумма
package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
