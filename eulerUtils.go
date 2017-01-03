package euler

import (
	"strconv"
	"math"
)


func GetFibonacci( n int ) []int {
	fib := []int {1, 2};
	for true{
		newFib := fib[len(fib)-1] + fib[len(fib)-2];
		if newFib > n {
			return fib
		}
		fib = append(fib, newFib);
	}
	return fib;
}

func PrimeFactors( n uint64 ) []uint64 {
	dividers := []uint64{}
	for d := uint64(2); n > d; d++ {
		if n%d == 0 {
			dividers = append(dividers, d)
			dividers = append(dividers, PrimeFactors(n/d)...)
			break;
		}
	}
	if len(dividers) == 0 { // Prime number
		dividers = append(dividers, n)
	}
	return dividers
}

func IsPalindrome( n int ) bool {
	stringNum := strconv.Itoa(n);
	for i := 0; i < int(len(stringNum)/2); i++ {
		if stringNum[i] != stringNum[len(stringNum)-1-i] {
			return false;
		}
	}
	return true;
}

func GetPrimesUntil( n int ) []int {
	primes := []int{2}
	for i := 3; i < n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}

func GetFirstPrimes( n int ) []int {
	primes := []int{2}
	for i := 3; len(primes) < n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPrime( n int, primes []int) bool {
	maxPossibleFac := int(math.Sqrt(float64(n)))
	for _, p := range primes{
		if n%p == 0 {
			return false
		}
		if p > maxPossibleFac {
			return true
		}
	}
	return true;
}
