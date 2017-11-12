package prime

func Sieve(max int64) (primes []int64) {
	b := make(map[int64]bool, 1000)
	for i := int64(2); i < max; i++ {
		if b[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < max; k += i {
			b[k] = true
		}
	}
	return
}
