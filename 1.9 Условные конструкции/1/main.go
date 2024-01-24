package main

import "fmt"

const (
	ABOVE = "Число положительное"
	BELOW = "Число отрицательное"
	ZERO  = "Ноль"
)

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(CheckNumber(number))
}

func CheckNumber(number int) string {
	switch {
	case number > 0:
		return ABOVE
	case number < 0:
		return BELOW
	}
	return ZERO
}
