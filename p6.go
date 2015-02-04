/*
Problem 6 - Find the difference between the sum of the squares of the first one hundred natural numbers and the square of their sum.
			ie, (1+2+...+100)^2 - (1^2+2^2+...+100^2)
*/
package main

import "fmt"

func main() {
	var sum, sqsum int
	for i := 1; i <= 100; i++ {
		sum += i;
		sqsum += i * i;
	}
	sum = sum * sum;
	fmt.Println(sum - sqsum)
}