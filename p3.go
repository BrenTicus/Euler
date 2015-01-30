/*
Problem 3 - Find the largest prime factor of 600851475143.
*/

package main

import "fmt"

// This is where I thought "You know, it'll probably take me longer to figure out an efficient approach than it 
// will to run the naive one." I was probably wrong.

// Checks if a number is prime.
func isPrime(n int) bool {
	for i := 2; i < n / 2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var i, largest int
	target := 600851475143
	// Only really have to check half the numbers since anything larger than target/2 is guaranteed to not be a factor.
	// Also, work backwards so we can just grab the biggest one and skip anything lower.
	for i = target / 2; i > 0; i-- {
		if target % i == 0 {
			if isPrime(i) {
				largest = i
				break
			}
		}
	}
	fmt.Println(largest)
}