// Церемония взвешивания
package main

import "fmt"

func main() {
	var answer string = "Легкий вес"
	var weight int
	fmt.Scan(&weight)

	if weight >= 60 && weight < 64 {
		answer = "Первый полусредний вес"
	} else if weight >= 64 && weight < 69 {
		answer = "Полусредний вес"
	}

	fmt.Println(answer)
}
