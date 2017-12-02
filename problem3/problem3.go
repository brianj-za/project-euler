package main

import (
	"fmt"

	"bitbucket.org/brianj-za/project-euler/prime"
	"bitbucket.org/brianj-za/project-euler/slice"
)

func main() {
	primes := slice.Max(prime.PrimeFactors(600851475143))
	fmt.Println(primes)
}
