// Любит - не любит
package main

import "fmt"

func main() {
	var (
		answer string = "любит"
		number int
	)

	fmt.Scan(&number)

	if number&1 == 1 {
		answer = "не " + answer
	}

	fmt.Println(answer)
}
