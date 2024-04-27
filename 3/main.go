package main

import "fmt"

func main() {
	fmt.Println(getPrimes(11, 20))
}

func getPrimes(min, max int) []int {
	var primes []int

	for i := min; i <= max; i++ {
		isPrime := true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
