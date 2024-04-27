package main

import (
	"fmt"
	"strconv"
)

func main() {
	printMultiplicationTable(5)
}

func printMultiplicationTable(num int) {
	digitLen := len(strconv.Itoa(num * num))

	fmt.Printf("%*s", digitLen+2, "")
	for i := 1; i <= num; i++ {
		fmt.Printf("%*d ", digitLen, i)
	}
	fmt.Println()

	for i := 1; i <= num; i++ {
		fmt.Printf("%*d", digitLen+2, i)
		for j := 1; j <= num; j++ {
			fmt.Printf("%*d ", digitLen, i*j)
		}
		fmt.Println()
	}
}
