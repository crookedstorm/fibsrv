package fibber

import (
	"fmt"
	"math/big"
)

func (seq MainSequence) BigFib(nth int) (big.Int, string) {
	// LookupFib returns a value in the in-memory table of Fibonacci numbers
	// if nth < 95 {
	// 	x := big.NewInt(0)
	// 	return x, errors.New("use LookupFib")
	// }
	var prev big.Int
	var next big.Int
	a, _ := seq.LookupFib(92)
	b, _ := seq.LookupFib(93)
	prev.SetUint64(a)
	next.SetUint64(b)

	for i := 94; i <= nth; i++ {
		prev.Add(&prev, &next)
		next, prev = prev, next
	}
	return prev, fmt.Sprintf("got %d and %d", a, b)
}
