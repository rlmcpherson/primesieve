package main

import (
	"fmt"
	"testing"
)

func BenchmarkPrintPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenPrimes(1000)
	}

}

func TestSumPrimes1000(t *testing.T) {
	const psum = 3682913
	var calcsum int
	primes := GenPrimes(1000)
	for _, p := range primes {
		calcsum += p
	}
	fmt.Println(primes[len(primes)-1])
	if psum != calcsum {
		t.Errorf("expected %d sum of primes, got %d\n", psum, calcsum)
	}
}
