package slice

// Sum adds all the values of the slice together.
func Sum(numbers []int64) (total int64) {
	for _, i := range numbers {
		total += i
	}
	return
}

// Max returns the largest value in the range of numbers
func Max(numbers []int64) (max int64) {
	for _, x := range numbers {
		if max < x {
			max = x
		}
	}
	return
}
