package euler

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func GetFibonacci(n int) []int {
	fib := []int{1, 2}
	for true {
		newFib := fib[len(fib)-1] + fib[len(fib)-2]
		if newFib > n {
			return fib
		}
		fib = append(fib, newFib)
	}
	return fib
}

func PrimeFactors(n uint64) []uint64 {
	dividers := []uint64{}
	for d := uint64(2); n > d; d++ {
		if n%d == 0 {
			dividers = append(dividers, d)
			dividers = append(dividers, PrimeFactors(n/d)...)
			break
		}
	}
	if len(dividers) == 0 { // Prime number
		dividers = append(dividers, n)
	}
	return dividers
}

func IsPalindrome(n int) bool {
	stringNum := strconv.Itoa(n)
	for i := 0; i < int(len(stringNum)/2); i++ {
		if stringNum[i] != stringNum[len(stringNum)-1-i] {
			return false
		}
	}
	return true
}

func GetPrimesUntil(n int) []int {
	primes := []int{2}
	for i := 3; i < n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}

func GetCollatzSequenceLength(n int) int {
	seqLength := 1
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		seqLength++
	}
	return seqLength
}

func GetFirstPrimes(n int) []int {
	primes := []int{2}
	for i := 3; len(primes) < n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPrime(n int, primes []int) bool {
	maxPossibleFac := int(math.Sqrt(float64(n)))
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
		if p > maxPossibleFac {
			return true
		}
	}
	return true
}

func ReadLinesFrom(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
