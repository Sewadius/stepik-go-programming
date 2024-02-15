package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)

	for n > 0 {
		var number int
		fmt.Scan(&number)

		if number > 9 && number < 100 {
			if number%8 == 0 {
				sum += number
			}
		}
		n--
	}
	fmt.Println(sum)
}
