package prime

import "math"

func PrimeFactors(number int64) []int64 {
	primes := Sieve(int64(math.Sqrt(float64(number))))

	factors := make([]int64, 0, 20)

	for _, n := range primes {
		if number%n == 0 {
			factors = append(factors, n)
		}
	}
	return factors
}
