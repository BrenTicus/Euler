/*
Problem 9 - Find the sum a + b + c = 1000 such that a^2 + b^2 = c^2.
*/
package main

import "fmt"

func main() {
	for c:=0; c < 1000; c++ {
		for b:=0; b < c; b++ {	// We can comfortably assume c > b
			for a:=0; a < b; a++ { // We also know c > a, and a != b due to the nature of pythagorean triples, so lock in that a < b
				// We could determine better bounds for b and a, but it doesn't matter for such a small set.
				if(a + b + c == 1000) {
					if(a * a + b * b == c * c) {
						fmt.Println(a * b * c);
						return;		// We know there's only one triple in this range, so just quit here.
					}
				}
			}
		}
	}
}