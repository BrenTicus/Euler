/*
Problem 4 - Find the largest number N such that N is a palindrome and the product of two 3-digit numbers.
*/

package main

import "fmt"
import "strconv"

// Reverses a string.
func reverse(in string) string {
	s := []rune(in)
	var i int = 0
	var j int = len(in)-1
	for ; i < len(s)/2; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}

func main() {
	var largest int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			// Convert to string and check if it's a palindrome.
			f := strconv.Itoa(i * j)
			if f == reverse(f) {
				if i * j > largest {
					// If it's a palindrome and the biggest yet, save it.
					largest = i * j
				}
			}
		}
	}
	fmt.Println(largest)
}