// YES or NO вот в чем вопрос
package main

import "fmt"

func main() {
	const (
		YES = "YES"
		NO  = "NO"
	)

	var n int
	fmt.Scan(&n)

	if n%2 != 0 {
		fmt.Println(YES)
	} else if n%2 == 0 && 2 <= n && n <= 5 {
		fmt.Println(NO)
	} else if n%2 == 0 && 6 <= n && n <= 20 {
		fmt.Println(YES)
	} else if n%2 == 0 && n > 20 {
		fmt.Println(NO)
	}
}
