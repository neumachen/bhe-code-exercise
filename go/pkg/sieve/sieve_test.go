package sieve

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime_Fail(t *testing.T) {
	t.Parallel()

	sieve := NewSieve()

	actual, err := sieve.NthPrime(-1)
	assert.Error(t, err)
	assert.Empty(t, actual)
}

func TestNthPrime_Success(t *testing.T) {
	t.Parallel()

	testData := []struct {
		nthPrime int64
		expected int64
	}{
		{
			nthPrime: 0,
			expected: 2,
		},
		{
			nthPrime: 1,
			expected: 3,
		},
		{
			nthPrime: 2,
			expected: 5,
		},
		{
			nthPrime: 5,
			expected: 13,
		},
		{
			nthPrime: 10,
			expected: 31,
		},
		{
			nthPrime: 19,
			expected: 71,
		},
		{
			nthPrime: 25,
			expected: 101,
		},
		{
			nthPrime: 38,
			expected: 167,
		},
		{
			nthPrime: 44,
			expected: 197,
		},
		{
			nthPrime: 52,
			expected: 241,
		},
		{
			nthPrime: 66,
			expected: 331,
		},
		{
			nthPrime: 99,
			expected: 541,
		},
		{
			nthPrime: 150,
			expected: 877,
		},
		{
			nthPrime: 200,
			expected: 1229,
		},
		{
			nthPrime: 500,
			expected: 3581,
		},
		{
			nthPrime: 800,
			expected: 6143,
		},
		{
			nthPrime: 986,
			expected: 7793,
		},
		{
			nthPrime: 2000,
			expected: 17393,
		},
		{
			nthPrime: 5000,
			expected: 48619,
		},
		{
			nthPrime: 10000,
			expected: 104743,
		},
		{
			nthPrime: 100000,
			expected: 1299721,
		},
		{
			nthPrime: 1000000,
			expected: 15485867,
		},
	}
	sieve := NewSieve()

	for i := range testData {
		tData := testData[i]
		actual, err := sieve.NthPrime(tData.nthPrime)
		assert.NoError(t, err)
		assert.Equal(t, tData.expected, actual, fmt.Sprintf("nth prime: %d", tData.nthPrime))
	}
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		t.Parallel()
		actual, err := sieve.NthPrime(n)
		assert.NoError(t, err)
		if !big.NewInt(actual).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
