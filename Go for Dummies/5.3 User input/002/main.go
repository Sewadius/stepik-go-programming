package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите предложение: ")

	input, _ := reader.ReadString('\n')
	fmt.Print("Вы ввели: ", input)
}
