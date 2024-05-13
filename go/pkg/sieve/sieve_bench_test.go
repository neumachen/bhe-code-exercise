package sieve

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkNthPrime benchmarks the NthPrime function across various ranges of n.
func BenchmarkNthPrime(b *testing.B) {
	// Cases for benchmarking different sizes of n.
	testCases := []struct {
		name string
		nth  int64
	}{
		{"10th Prime", 10},
		{"100th Prime", 100},
		{"1000th Prime", 1000},
		{"10000th Prime", 10000},
		{"100000th Prime", 100000},
		{"1000000th Prime", 1000000},
		{"10000000th Prime", 10000000},
		{"100000000th Prime", 100000000},
	}

	// Initialize the Sieve once to use for all benchmark cases.
	sieve := NewSieve()

	for i := range testCases {
		testCase := testCases[i]
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := sieve.NthPrime(testCase.nth)
				assert.NoError(
					b,
					err,
					fmt.Sprintf(
						"failed to calculate prime for %s, error: %s", testCase.name, err,
					),
				)
			}
		})
	}
}
