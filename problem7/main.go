package main

import (
	"fmt"

	"bitbucket.com/brianj-za/project-euler/prime"
)

func main() {
	primes := prime.SieveByCount(10001)
	fmt.Println(primes[len(primes)-1])
}
