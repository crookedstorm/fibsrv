// Package fibber provides a type and methods for quickly returning the nth number of the
// Fibonacci sequence.
package fibber

import "errors"

// MainSequence is an array set to contain the Fibonacci numbers up to the 94th.
// Doing a simple array lookup is much faster than calculating numbers where possible.
type MainSequence [94]uint64

// InitCache builds the in-memory Fibonacci sequence to reference
// up to the 94th number (the largest uint64 possible)
func (seq *MainSequence) InitCache() error {
	for i := 0; i < len(seq); i++ {
		if i == 0 {
			seq[i] = 0
			continue
		} else if i == 1 {
			seq[i] = 1
			continue
		}
		seq[i] = seq[i-1] + seq[i-2]
	}
	return nil
}

// LookupFib returns a value in the in-memory table of Fibonacci numbers
func (seq MainSequence) LookupFib(nth int) (uint64, error) {
	if nth > 94 {
		return 0, errors.New("cannot lookup greater than 94th number")
	}
	return seq[nth], nil
}
