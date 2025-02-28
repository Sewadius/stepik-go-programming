// Какой сегодня месяц?
package main

import "fmt"

func main() {
	var result string = "Некорректный номер месяца"
	var month uint8
	fmt.Scan(&month)

	switch month {
	case 1:
		result = "Январь"
	case 2:
		result = "Февраль"
	case 3:
		result = "Март"
	case 4:
		result = "Апрель"
	case 5:
		result = "Май"
	case 6:
		result = "Июнь"
	case 7:
		result = "Июль"
	case 8:
		result = "Август"
	case 9:
		result = "Сентябрь"
	case 10:
		result = "Октябрь"
	case 11:
		result = "Ноябрь"
	case 12:
		result = "Декабрь"
	}

	fmt.Println(result)
}
