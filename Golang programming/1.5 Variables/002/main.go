package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(GetResult(a, b))
}

func GetResult(a, b int) int {
	return a*a + b*b
}
