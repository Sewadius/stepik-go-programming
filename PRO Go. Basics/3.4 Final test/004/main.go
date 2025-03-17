// Возрастающая последовательность
package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	a := number / 100
	b := number / 10 % 10
	c := number % 10

	if a < b && b < c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
