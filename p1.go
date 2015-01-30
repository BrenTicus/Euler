/*
Problem 1 - Find the sum of all multiples of 3 and 5 less than 1000.
*/

import "fmt"

func main() {
	var sum int
	for i := 1; i < 1000; i++ {
		if i % 3 == 0 {
			sum = sum + i
		} else if i % 5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}