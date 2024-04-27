package main

import "fmt"

func main() {
	println(computers(25))
	println(computers(41))
	println(computers(1048))
}

func computers(num uint64) string {
	var res string
	lastDigit := num % 10

	if lastDigit == 1 {
		res = "1 компьютер"
	} else if 2 <= lastDigit && lastDigit <= 4 {
		res = fmt.Sprintf("%d компьютера", num)
	} else {
		res = fmt.Sprintf("%d компьютеров", num)
	}

	return res
}
