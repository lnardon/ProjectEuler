package main

import (
	"fmt"
	"math"
)

// The prime factors of $13195$ are $5, 7, 13$ and $29$.
// What is the largest prime factor of the number $600851475143$
// This problem has a difficulty rating of 5%. 
func main(){
   largest := get_largest_prime_factor(600851475143)
   fmt.Println("Largest prime factor is: %i" , largest)
}

func get_largest_prime_factor(n int64) int64 {
	var lastPrime int64 = 0
	var currentPrime int64 = 2
	var factoredValue = n

	for factoredValue > 1 {
		if factoredValue % currentPrime == 0 {
			factoredValue /= currentPrime
			lastPrime = currentPrime
		} else {
			currentPrime = get_next_prime(currentPrime)
		}
	}

	return lastPrime
}

func get_next_prime(last int64) int64 {
	next := last + 1
	for {
		isPrime := true
		for i := int64(2); i <= int64(math.Sqrt(float64(next))); i++ {
			if next%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return next
		}
		next++
	}
}
