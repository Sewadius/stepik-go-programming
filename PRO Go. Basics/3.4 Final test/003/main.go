package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	one := n / 1000
	two := n / 100 % 10
	three := n / 10 % 10
	four := n % 10

	var check bool = one == four && two == three

	if check {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
