package main

import (
	"fmt"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is $9009 = 91 \times 99$.
// Find the largest palindrome made from the product of two $3$-digit numbers.
// This problem has a difficulty rating of 5%
func main() {
	largestPalindrome := 0
	var n1, n2 int

	for first_number := 999; first_number >= 100; first_number-- {
		for second_number := first_number; second_number >= 100; second_number-- {
			product := first_number * second_number
			if isPalindrome(product) && product > largestPalindrome {
				largestPalindrome = product
				n1 = first_number
				n2 = second_number
			}
		}
	}

	fmt.Printf("Largest palindrome is %d, from %d and %d\n", largestPalindrome, n1, n2)
}

func isPalindrome(n int) bool {
	reverse := 0
	temp := n
	for temp > 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}
	return reverse == n
}

