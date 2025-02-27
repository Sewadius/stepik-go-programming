// Длина слова
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var word string
	fmt.Scan(&word)

	inBytes := len(word)
	inRunes := utf8.RuneCountInString(word)

	fmt.Println(inBytes, inRunes)
}
