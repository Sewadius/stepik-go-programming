// Сложность пароля
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var password string
	fmt.Scan(&password)

	length := utf8.RuneCountInString(password)

	if length < 9 {
		fmt.Println("Легкий")
	} else if length < 13 {
		fmt.Println("Средний")
	} else if length < 17 {
		fmt.Println("Сложный")
	} else {
		fmt.Println("Очень сложный")
	}
}
