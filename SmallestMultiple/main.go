package main 

import ("fmt")

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible divisible with no remainder by all of the numbers from 1 to 20?
// This problem has a difficulty rating of 5%.
func main(){
	current_number := 1

	for i := 1; i <= 20 ; i++ {
		if current_number % i != 0 {
			current_number += 1
			i = 1
		}
	}

	fmt.Printf("Smallest multiple is: %d \n", current_number)

}