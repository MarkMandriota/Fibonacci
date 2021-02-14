package fib

//go:noinline
// Fib - returns the fibonacci number by ordinal
func Fib(n uint64) uint64

func New() func() uint64 {
	var x, y uint64 = 0, 1

	return func() uint64 {
		x, y = x+y, x
		return y
	}
}
