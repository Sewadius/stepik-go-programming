// Возрастная группа
package main

import "fmt"

func main() {
	const S string = "разработчик"
	var n int
	fmt.Scan(&n)

	if n <= 3 {
		fmt.Println("начинающий")
	} else if n <= 7 {
		fmt.Println("младший", S)
	} else if n <= 15 {
		fmt.Println("средний", S)
	} else {
		fmt.Println("старший", S)
	}
}
