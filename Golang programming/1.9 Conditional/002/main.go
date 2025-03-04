package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	first := number % 10
	second := number / 10 % 10
	third := number / 100

	if first == second || second == third || first == third {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
