// Координатная четверть
package main

import "fmt"

func main() {
	var x, y float32
	var answer int = 1

	fmt.Scan(&x, &y)

	if x < 0 && y > 0 {
		answer = 2
	} else if x < 0 && y < 0 {
		answer = 3
	} else if x > 0 && y < 0 {
		answer = 4
	}

	fmt.Println(answer)
}
