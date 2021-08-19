package fibber

import "errors"

type MainSequence [94]uint64

func (seq *MainSequence) InitCache() error {
	// InitCache builds the in-memory Fibonacci sequence to reference
	// up to the 94th number (the largest uint64 possible)
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

func (seq MainSequence) LookupFib(nth int) (uint64, error) {
	// LookupFib returns a value in the in-memory table of Fibonacci numbers
	if nth > 94 {
		return 0, errors.New("cannot lookup greater than 94th number")
	}
	return seq[nth], nil
}
