// Количество равных из трех
package main

import "fmt"

func main() {
	var result int = 0
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c && a == c {
		result = 3
	} else if a == b || b == c || a == c {
		result = 2
	}

	fmt.Println(result)
}
