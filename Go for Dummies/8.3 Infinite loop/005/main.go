// Первое простое в бесконечном потоке
package main

import "fmt"

func main() {
	var number int
	var isFond bool

	for {
		fmt.Scan(&number)

		for i := 2; i <= number; i++ {
			if number%i == 0 && number != i && number > 2 {
				break
			}
			if number == i {
				isFond = true
			}
		}

		if isFond {
			break
		}
	}

	fmt.Println(number)
}
