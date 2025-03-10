// Принадлежность 3
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var check_1 bool = -30 < x && x <= -2
	var check_2 bool = 7 < x && x <= 25

	if check_1 || check_2 {
		fmt.Println("Принадлежит")
	} else {
		fmt.Println("Не принадлежит")
	}
}
