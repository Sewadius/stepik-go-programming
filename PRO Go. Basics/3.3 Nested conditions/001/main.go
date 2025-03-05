// Принадлежность точки
package main

import "fmt"

func main() {
	var (
		answer string = "NO"
		x      int
	)

	fmt.Scan(&x)

	var check_1 bool = x >= -3 && x <= 1
	var check_2 bool = x >= 5 && x <= 9

	if check_1 || check_2 {
		answer = "YES"
	}

	fmt.Println(answer)
}
