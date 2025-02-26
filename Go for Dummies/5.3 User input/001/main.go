package main

import "fmt"

func main() {
	var word1, word2 string

	fmt.Print("Введите два слова, разделенные пробелом: ")

	fmt.Scan(&word1, &word2)
	fmt.Println("Вы ввели:", word1, word2)
}
