package main

import ("fmt")

// The following iterative sequence is defined for the set of positive integers:
// n -> n/2 (n is even)
// n -> 3n + 1 (n is odd)
// Using the rule above and starting with 13, we generate the following sequence:
// 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1.
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
// Which starting number, under one million, produces the longest chain?
// NOTE: Once the chain starts the terms are allowed to go above one million.
func main(){
	lgst := &Longest{
		number : 1,
		size: 1,
	}

	for i := 1; i < 1000000; i++ {
        aux := get_sequence(i)
        if aux.size > lgst.size {
            lgst = aux
        }
    }

    fmt.Printf("Number: %d has the chain size of %d\n", lgst.number, lgst.size)
}

type Longest struct {
	number int
	size   int
}

func get_sequence(seed int) *Longest {
    current := seed
    size := 1

    for {
        if current % 2 == 0 {
            current = current / 2
        } else {
            current = 3 * current + 1
        }

        size++
        if current == 1 {
            return &Longest{
                number: seed,
                size:   size,
            }
        }
    }
}