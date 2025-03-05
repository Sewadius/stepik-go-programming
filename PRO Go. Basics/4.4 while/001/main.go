// Сколько раз делится
package main

import "fmt"

func main() {
	const Three = 3
	var a uint

	fmt.Scan(&a)

	var counter uint8

	for a%Three == 0 {
		a /= Three
		counter++
	}

	fmt.Println(counter)
}
