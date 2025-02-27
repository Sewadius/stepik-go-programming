// Unicode триплет
package main

import "fmt"

func main() {
	var code1, code2, code3 rune
	fmt.Scan(&code1, &code2, &code3)

	result := string(code1) + string(code2) + string(code3)
	fmt.Println(result)
}
