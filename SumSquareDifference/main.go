package main 

import (
	"fmt"
)

// The sum of the squares of the first ten natural numbers is,
// $$1^2 + 2^2 + ... + 10^2 = 385.$$
// The square of the sum of the first ten natural numbers is,
// $$(1 + 2 + ... + 10)^2 = 55^2 = 3025.$$
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func main(){
	var square_of_the_sum,sum_of_squares int

	for i:= 1; i <= 100; i++ {
		sum_of_squares += (i * i)
		square_of_the_sum += i
	}
	square_of_the_sum *= square_of_the_sum

	fmt.Printf("The difference is: %d \n", square_of_the_sum - sum_of_squares)
}