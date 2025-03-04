package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	first := n / 100
	second := n / 10 % 10
	third := n % 10

	number := third*100 + second*10 + first
	fmt.Println(number)
}
