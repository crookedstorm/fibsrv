package fibber

import (
	"errors"
	"math/big"
)

// BigFib calculates Fibonacci numbers bigger than the 94th number.
// You cannot go higher than the 94th with uint64, so this is needed.
func (seq MainSequence) BigFib(nth int) (big.Int, error) {
	if nth < 95 {
		x := big.NewInt(0)
		return *x, errors.New("use LookupFib")
	}
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
	return prev, nil
}
