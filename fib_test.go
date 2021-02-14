package fib

import (
	"testing"
)

func Test_Fib(t *testing.T) {
	fib := New()
	for i := uint64(0); i < 100; i++ {
		var result uint64

		t.Run("Fibonacci", func(t *testing.T) {
			result = fib()
		})

		t.Logf("Fibonacci(%d) = %d\n", i, result)
	}
}
