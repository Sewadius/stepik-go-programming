package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for n > 9 {
		n /= 10
	}

	fmt.Println(n)
}
