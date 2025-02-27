// Дневник
package main

import "fmt"

func main() {
	const (
		VeryBad      = 1
		Bad          = 2
		Satistactory = 3
		Good         = 4
		Excellent    = 5
	)

	var mark uint8
	fmt.Scan(&mark)

	switch mark {
	case VeryBad:
		fmt.Println("Очень плохо")
	case Bad:
		fmt.Println("Плохо")
	case Satistactory:
		fmt.Println("Удовлетворительно")
	case Good:
		fmt.Println("Хорошо")
	case Excellent:
		fmt.Println("Отлично")
	default:
		fmt.Println("Некорректная оценка")
	}
}
