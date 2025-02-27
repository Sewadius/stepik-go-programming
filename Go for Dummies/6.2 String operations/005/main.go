// Символьная трансформация
package main

import "fmt"

func main() {
	var code1, code2 rune
	fmt.Scan(&code1, &code2)

	result := code1 + code2
	fmt.Println(string(result))
}
