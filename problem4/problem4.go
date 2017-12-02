package main

import (
	"log"
	"time"
)

func isPalindrome(n int) bool {
	if (n/100000) == n%10 &&
		(n/10000)%10 == (n/10)%10 &&
		(n/1000)%10 == (n/100)%10 {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	var highest int = 0

	for i := 999; i > 100; i-- {
		dec := 1
		if i%11 == 0 {
			dec = 11
		}
		for j := 999; j > 100; j -= dec {
			p := i * j
			if p > highest && isPalindrome(p) {
				highest = p
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Highest %d (took: %s)\n", highest, elapsed)
}
