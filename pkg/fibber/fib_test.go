package fibber

import (
	"os"
	"testing"
)

var f MainSequence

func TestMain(m *testing.M) {
	f.InitCache()
	code := m.Run()
	os.Exit(code)
}
func TestCalculatesCorrectFibonaccis(t *testing.T) {
	ninth, err := f.LookupFib(9)
	if ninth != 34 {
		t.Error("Initial Fibonacci cache isn't correct")
	}
	if err != nil {
		t.Errorf("Something is wrong: %v\n", err)
	}
}

func TestBigFibonaccis(t *testing.T) {
	bigone, err := f.BigFib(300)
	if bigone.String() != "137347080577163115432025771710279131845700275212767467264610201" {
		t.Errorf("Huge Fibonacci number isn't correct, got %v\n", bigone)
	}
	if err != nil {
		t.Errorf("Something is wrong: %v\n", err)
	}
}

func benchmarkLargeFibonacci(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f.BigFib(220)
	}
}

func benchmarkFibLookups(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f.LookupFib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFibLookups(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFibLookups(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFibLookups(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFibLookups(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFibLookups(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFibLookups(40, b) }

func BenchmarkBigFib1(b *testing.B)  { benchmarkLargeFibonacci(95, b) }
func BenchmarkBigFib2(b *testing.B)  { benchmarkLargeFibonacci(96, b) }
func BenchmarkBigFib3(b *testing.B)  { benchmarkLargeFibonacci(97, b) }
func BenchmarkBigFib10(b *testing.B) { benchmarkLargeFibonacci(110, b) }
func BenchmarkBigFib20(b *testing.B) { benchmarkLargeFibonacci(200, b) }
func BenchmarkBigFib40(b *testing.B) { benchmarkLargeFibonacci(400, b) }
