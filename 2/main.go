package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(commonDivisors([]int{42, 12, 18}))
}

func commonDivisors(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	min := slices.Min(nums)

	var commonDivisors []int

	for i := 2; i <= min; i++ {
		isCommonDivisior := true

		for _, num := range nums {
			if num%i != 0 {
				isCommonDivisior = false
				break
			}
		}

		if isCommonDivisior {
			commonDivisors = append(commonDivisors, i)
		}
	}

	return commonDivisors
}
