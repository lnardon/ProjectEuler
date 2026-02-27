package main

import (
	"fmt"
	"euler/helpers"
)

func main(){
	defer helpers.MeasureTimeSpent()()

	var sum int = 0;
	for i := 1; i < 2000000; i++ {
		if helpers.IsPrime(i) {
			sum += i
		}
	}
	fmt.Printf("The sum of all the primes below two million is: %d\n", sum)
}
