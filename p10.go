/*
Problem 9 - Find the sum of all primes under 2,000,000.
*/

package main

import "fmt"

// Checks if a number is prime.
// Should really figure out a better way to do this...
func isPrime(n int) bool {
	if n % 2 == 0 {		
		return false
	}
	for i := 3; i < n / 2; {
		if n % i == 0 {
			return false
		}
		i += 2
	}
	return true
}

func main() {
	sum := 2;		//We'll just start here so we can start the prime count at 3.
	for i:=3; i < 2000000; {
		if(isPrime(i)) {
			sum += i
		}
		i+=2
	}
	fmt.Println(sum)
}