// Спасите неопытного Джуна
package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Конец программы")
}
