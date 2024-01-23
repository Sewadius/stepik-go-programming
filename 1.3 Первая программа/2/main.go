package main

import "fmt"

func main() {
	for i := range make([]bool, 3) {
		fmt.Println("I like Go!")
		i++
		i--
	}
}
