package main

import "fmt"

func main() {
	var (
		word1 string = "Hello"
		word2 string = "World"
	)

	wordsTogether := word1 + word2
	fmt.Println(wordsTogether) // "HelloWorld"
}
