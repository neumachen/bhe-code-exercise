package sieve

import (
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
			nthPrime: 19,
			expected: 71,
		},
		{
			nthPrime: 99,
			expected: 541,
		},
		{
			nthPrime: 500,
			expected: 3581,
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
			nthPrime: 1000000,
			expected: 15485867,
		},
	}
	sieve := NewSieve()

	for i := range testData {
		tData := testData[i]
		actual, err := sieve.NthPrime(tData.nthPrime)
		assert.NoError(t, err)
		assert.Equal(t, tData.expected, actual)
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
