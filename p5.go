/*
Problem 5 - Find the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.
*/

package main

import "fmt"

func main() {
	i := 2520	// The problem gives 2520 as the smallest number divisible by 1 to 10, so it seems like a good start point.
	for ; ; i = i + 20 {	// We can just skip straight to numbers divisible by 20.
		// This is gross, but tests everything. Smaller numbers are implied by these (ie, if 14|i then 7|i)
		if(i % 19 == 0 && i % 18 == 0 && i % 17 == 0 && i % 16 == 0 && i % 15 == 0 && 
			i % 14 == 0 && i % 13 == 0 && i % 12 == 0 && i % 11 == 0) {
			break	// Yeah, we can just stop. Seems pointless looking at it like this.
		}
	}
	fmt.Println(i);
}