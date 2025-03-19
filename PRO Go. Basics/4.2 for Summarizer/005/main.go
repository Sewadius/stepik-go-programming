// Нули
package main

import "fmt"

func main() {
	var n, number int
	var answer string = "NO"

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number == 0 {
			answer = "YES"
		}
	}

	fmt.Println(answer)
}
