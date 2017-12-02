package main

import (
	"fmt"

	"bitbucket.org/brianj-za/project-euler/slice"
)

func main() {
	var i int64 = 1
	numbers := make([]int64, 0, 100)
	for {
		if i%3 == 0 || i%5 == 0 {
			numbers = append(numbers, i)
		}
		i++
		if i >= 1000 {
			break
		}
	}
	fmt.Printf("Total: %d\n", slice.Sum(numbers))
}
