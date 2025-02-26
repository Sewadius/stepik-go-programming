package main

import "fmt"

func main() {
	var a string = "a"
	var b string

	b = "b"
	c := "c"

	var d string = a
	e := c

	f := b
	fmt.Println(a, b, c, d, e, f) // a b c a c b
}
