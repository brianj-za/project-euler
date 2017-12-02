package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	m := 20
	found := true
	val := 0

	for i := m; ; i += m {

		for j := m - 1; j > 1; j-- {
			if i%j != 0 {
				found = false
				break
			}

		}
		if found {
			val = i
			break
		}
		found = true

	}
	log.Printf("Found: %d\n", val)

	elapsed := time.Since(start)
	log.Printf("Elapsed: %s\n", elapsed)
}
