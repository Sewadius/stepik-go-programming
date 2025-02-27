// Конкатенация троих
package main

import "fmt"

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)

	result := s1 + s2 + s3
	fmt.Println(result)
}
