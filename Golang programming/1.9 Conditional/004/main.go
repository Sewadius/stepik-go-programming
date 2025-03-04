// Счастливое число
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	six := n % 10
	five := n / 10 % 10
	four := n / 100 % 10

	one := n / 100000
	two := n / 10000 % 10
	three := n / 1000 % 10

	sum1 := one + two + three
	sum2 := six + five + four

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
