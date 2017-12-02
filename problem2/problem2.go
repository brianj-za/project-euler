package main

import (
	"fmt"

	"bitbucket.org/brianj-za/project-euler/fibonacci"
	"bitbucket.org/brianj-za/project-euler/slice"
)

func main() {
	f := fibonacci.Sequence()

	numbers := make([]int64, 0, 20)
	for x := int64(0); x < 4000000; x = f() {
		if x&1 == 0 {
			numbers = append(numbers, x)
		}
	}
	fmt.Println(slice.Sum(numbers))
}
