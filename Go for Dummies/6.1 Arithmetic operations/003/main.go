package main

import "fmt"

func main() {
	var a, b, c float32

	fmt.Scan(&a, &b, &c)

	p1 := 234 * b * c * (5 + 12*c) / 259
	p2 := 22*a - p1*a*c

	fmt.Println(int(p2))
}
