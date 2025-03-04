package main

import (
	"fmt"
	"strconv"
)

const ONE_HOUR = 30

func main() {
	var degrees int
	fmt.Scan(&degrees)
	fmt.Println(GetResult(degrees))
}

func GetResult(d int) string {
	hours := strconv.FormatInt(int64(d/ONE_HOUR), 10)
	minutes := strconv.FormatInt(int64(2*(d%ONE_HOUR)), 10)
	return "It is " + hours + " hours " + minutes + " minutes."
}
