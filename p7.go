/*
Problem 7 - Find the 10001st prime.
*/
package main

import "fmt"

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
	primeCount:=6
	big:=1
	for i:=17; primeCount < 10001; {	//Problem gives 13 as the 6th prime; 15 isn't prime; just start at 17.
		if(isPrime(i)) {
			primeCount++
			big = i
		}
		i+=2
	}
	fmt.Println(big)
}