package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "Этот текст состоит из 74 байт и 43 символов"

	var lengthInBytes int = len(text)
	var lengthInRunes int = utf8.RuneCountInString(text)

	fmt.Println(lengthInBytes, lengthInRunes)
}
