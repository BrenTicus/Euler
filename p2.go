/*
Problem 2 - Find the sum of all even fibonacci numbers less than 4000000.
*/

package main

import "fmt"

func main() {
	twoback := 1
	oneback := 2
	sum := 2
	// Just start at three and initialize values to the other variables.
	for i := 3; i < 4000000 ; {
		twoback = oneback
		oneback = i
		i = oneback + twoback
		if i % 2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}