package main

import "fmt"

func main() {
	var n, number, sum int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number%2 == 0 && number%3 != 0 {
			sum += number
		}
	}

	fmt.Println(sum)
}
