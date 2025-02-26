package main

import "fmt"

func main() {
	var someText string = ""
	fmt.Println("some text: ", someText)

	someText = "new text"
	fmt.Println("some text: ", someText)

	moreText := someText
	fmt.Println("more text: ", moreText)
}
