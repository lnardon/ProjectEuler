package main

import "euler/helpers"

func main() {
	defer helpers.MeasureTimeSpent()()

	var count int = 0
	var prime int = 2
	var number int = 0
	for count < 10001 {
		if isPrime(prime) {
			count++
			number = prime
		}
		prime++
	}
	println(number)
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
