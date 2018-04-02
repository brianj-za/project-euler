package main

import (
	"testing"

	"bitbucket.com/brianj-za/project-euler/prime"
)

func BenchmarkSieve(b *testing.B) {
	prime.SieveByCount(int64(b.N))
}
