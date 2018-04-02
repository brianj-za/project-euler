package prime

func Sieve(max int64) (primes []int64) {
	primes = make([]int64, 0, 100)
	ch := make(chan int64) // Create a new channel.
	go generate(ch)        // Launch Generate goroutine.
	var pch int64
	for i := int64(0); pch < max; i++ {
		pch = <-ch
		primes = append(primes, pch)
		// print(pch, "\n")
		ch1 := make(chan int64)
		go filter(ch, ch1, pch)
		ch = ch1
	}
	return
}

func SieveByCount(count int64) (primes []int64) {
	primes = make([]int64, 0, count)
	ch := make(chan int64) // Create a new channel.
	go generate(ch)        // Launch Generate goroutine.
	var pch int64
	for i := int64(0); i < count; i++ {
		pch = <-ch
		primes = append(primes, pch)
		// print(pch, "\n")
		ch1 := make(chan int64)
		go filter(ch, ch1, pch)
		ch = ch1
	}
	return
}

func generate(ch chan<- int64) {
	for i := int64(2); ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func filter(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
