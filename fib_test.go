package bencher

import "testing"

func BenchmarkFibonacci5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(5)
	}
}

func BenchmarkFibonacci10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(10)
	}
}
