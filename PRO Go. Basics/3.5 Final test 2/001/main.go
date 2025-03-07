// Играем в футбол
package main

import "fmt"

func main() {
	var (
		age    int
		gender string
		answer string = "NO"
	)

	fmt.Scan(&age, &gender)

	if age >= 12 && age <= 18 && gender == "m" {
		answer = "YES"
	}

	fmt.Println(answer)
}
