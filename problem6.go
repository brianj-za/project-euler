package main

import (
	"log"
	"time"

	"bitbucket.org/brianj-za/project-euler/slice"
)

func main() {
	start := time.Now()

	s := make([]int64, 100)
	for i := 0; i < 100; i++ {
		s[i] = int64(i + 1)
	}

	sum := slice.Sum(s)
	sumSquare := slice.SumSquares(s)

	elapsed := time.Since(start)
	log.Printf("Value: %d\n", (sum*sum)-sumSquare)
	log.Printf("Elapsed: %s\n", elapsed)

}
