package bencher

func Fibonacci(number int) int {
	a := 0
	b := 1
	// Iterate until desired position in sequence.
	for i := 0; i < number; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
	}
	return a
}
