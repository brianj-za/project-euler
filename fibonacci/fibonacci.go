package fibonacci

func Sequence() func() int64 {
	var a, b int64 = 1, 0
	return func() int64 {
		a, b = b, a+b
		return a
	}
}
