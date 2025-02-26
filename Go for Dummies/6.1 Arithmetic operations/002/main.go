// Спидометр
package main

import "fmt"

func main() {
	var (
		distance float32
		seconds  int
	)

	fmt.Scan(&distance, &seconds)

	speed := distance / float32(seconds)
	fmt.Println(speed)
}
