package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		fmt.Println("Hello, World!")
	}
}
