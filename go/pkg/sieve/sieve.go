package sieve

import (
	"fmt"
	"math"
)

// Sieve is an interface that defines the functionality for finding the nth prime number.
// It includes a single method, NthPrime, which takes an integer n and returns the nth prime number as an int64.
type Sieve interface {
	NthPrime(n int64) (int64, error)
}

// NewSieve creates and returns an instance of Sieve.
// This function initializes a siever with a specific prime-finding function, NthPrime in this case,
// encapsulating the behavior into a callable function object. This approach provides flexibility in swapping
// different prime calculation algorithms without altering the client code that uses this interface.
func NewSieve() Sieve {
	return siever(NthPrime)
}

// siever is a function type that takes an int64 and returns an int64, designed to calculate the nth prime number.
// It implements the Sieve interface, allowing function types to be used directly in places where Sieve is expected,
// promoting a functional programming approach within a structured interface.
type siever func(n int64) (int64, error)

// NthPrime is a method on the siever type that implements the Sieve interface's NthPrime method.
// It directly invokes the function represented by the siever instance to compute the nth prime number.
// This design leverages Go's first-class functions to dynamically specify the behavior of the NthPrime method at runtime,
// increasing the code's modularity and flexibility.
func (s siever) NthPrime(n int64) (int64, error) {
	return s(n)
}

// The following line is a compile-time assertion to ensure that the siever type correctly implements the Sieve interface.
// It serves as a safeguard, causing a compile-time error if the siever does not meet the Sieve interface requirements,
// thus providing early feedback and avoiding runtime issues related to interface implementation.
var _ Sieve = siever(NthPrime)

// sieve performs the Sieve of Eratosthenes algorithm to identify all prime numbers up to a specified limit.
// This function initializes an array to keep track of the prime status of each number starting from 2 up to 'limit'.
// It then marks the non-prime numbers based on the classic algorithm's approach:
// 1. Start from the first prime number, 2, mark all of its multiples as non-prime.
// 2. Move to the next number that is still marked as prime and repeat until you've processed numbers up to sqrt(limit).
// This method is efficient because it only requires O(n log log n) time complexity and O(n) space complexity, where
// n is the number of integers up to 'limit'.
func sieve(limit int) []int {
	isPrime := make([]bool, limit+1) // A slice that tracks the primality of numbers up to `limit`.
	for i := 2; i <= limit; i++ {
		isPrime[i] = true // Initialize all entries as true, indicating potential primes.
	}

	// Eliminate non-primes using the Sieve of Eratosthenes method.
	for p := 2; p*p <= limit; p++ {
		if isPrime[p] { // Only consider p as a prime if it has not been marked otherwise.
			for i := p * p; i <= limit; i += p {
				isPrime[i] = false // Mark multiples of p starting from p^2 as non-prime.
			}
		}
	}

	// Extract prime numbers from the sieve.
	var primes []int
	for i, prime := range isPrime {
		if prime {
			primes = append(primes, i) // Append the index of true entries to the result slice.
		}
	}

	return primes
}

// NthPrime calculates the nth prime number using an optimized sieve method based on the Prime Number Theorem. It
// dynamically adjusts the number range to sieve, ensuring efficient memory and time utilization. The function initially
// calculates an upper limit for the sieve using the Prime Number Theorem, which states that the nth prime is
// approximately n * (log n + log log n). This estimation ensures that the sieve operation is not wastefully large,
// balancing between performance and memory usage. If the calculated primes are fewer than required, the function
// returns an error indicating the estimation fell short. This ensures that any limitations of the Prime Number
// Theorem's approximation in practical scenarios are communicated back to the caller
func NthPrime(nth int64) (int64, error) {
	if nth < 1 {
		if nth == 0 {
			return 2, nil // short circuit when nth is 0
		}
		return 0, nil // There is no such thing as a zeroth prime, so return 0 for invalid input.
	}

	// original implementation
	// limit := 1000

	// Estimate the upper limit for the prime number calculation using the Prime Number Theorem.
	nthFloat := float64(nth)
	limit := int(nthFloat * (math.Log(nthFloat) + math.Log(math.Log(nthFloat))))
	primes := sieve(limit)
	if len(primes) < int(nth) {
		return 0, fmt.Errorf("prime number list generated up to %d does not contain %d primes", limit, nth)
	}
	return int64(primes[nth-1]), nil
}
