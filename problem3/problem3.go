package main

import (
	"fmt"

	"bitbucket.com/brianj-za/project-euler/prime"
	"bitbucket.com/brianj-za/project-euler/slice"
)

func main() {
	primes := slice.Max(prime.PrimeFactors(600851475143))
	fmt.Println(primes)
}
