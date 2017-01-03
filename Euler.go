package euler

import (
	"math"
	"fmt"
	"strconv"
	"time"
	"math/big"
)

func problem1() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if math.Mod(float64(i), 3) == 0 || math.Mod(float64(i), 5) == 0 {
			sum += i
		}
	}
	return sum
}

func problem2() int {
	fibs := GetFibonacci(4000000)
	sum := 0
	for _, f := range fibs {
		if f%2 == 0 {
			sum += f
		}
	}
	return sum
}

func problem3() int {
	n := uint64(600851475143)
	factors := PrimeFactors(n)
	return int(factors[len(factors)-1])
}

func problem4() int {
	for i := 999 * 999; i >= 0; i-- {
		if IsPalindrome(i) {
			for j := 999; j > 99; j-- {
				if i%j == 0 && i/j <= 999 {
					return i
				}
			}
		}
	}
	return 0
}

func problem5() int {
	var factors [20]int
	var maxFactors [20]int

	for i := 0; i < len(maxFactors); i++ {
		maxFactors[i] = 0
	}

	for i := 2; i < len(factors); i++ {
		for i := 0; i < len(maxFactors); i++ {
			factors[i] = 0
		}

		for _, f := range PrimeFactors(uint64(i)) {
			factors[f]++
		}

		for j, f := range factors {
			if f > maxFactors[j] {
				maxFactors[j] = f
			}
		}
	}

	sum := float64(1)
	for i, f := range maxFactors {
		if f != 0 {
			sum *= math.Pow(float64(i), float64(f))
		}
	}

	return int(sum)
}

func problem6() int {
	n := 100

	sumOfSquare := 0
	SquareOfSum := 0

	for i := 1; i <= n; i++ {
		sumOfSquare += i * i
		SquareOfSum += i
	}
	SquareOfSum = SquareOfSum * SquareOfSum

	return SquareOfSum - sumOfSquare
}

func problem7() int {
	loadsOfPrimes := GetFirstPrimes(10001)
	return loadsOfPrimes[len(loadsOfPrimes)-1]
}

func problem8() int {
	number := ReadLinesFrom("src/github.com/reicher/euler/data/prob8.txt")[0]
	maxProduct := 0
	cl := 13

	for i := 0; i <= len(number)-cl; i++ {
		product, _ := strconv.Atoi(string(number[i]))
		for j := i + 1; j < cl+i || j > len(number); j++ {
			num, _ := strconv.Atoi(string(number[j]))
			product *= num
		}

		if product > maxProduct {
			maxProduct = product
		}
		if (i + cl) >= len(number) {
			break
		}
	}

	return maxProduct
}

func problem9() int {
	for a := float64(1); a < 1000; a++ {
		for b := a + 1; b < 1000-a; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if a+b+c == 1000 && c > b {
				return int(a * b * c)
			} else if a+b+c > 1000 {
				break
			}
		}
	}
	return 0
}

func problem10() int {
	primes := GetPrimesUntil(2000000)
	sum := 0
	for _, p := range primes {
		sum += p
	}

	return sum
}

// Ugliest ever
func problem11() int {
	matrix := [20][20]int{[20]int{8, 02, 22, 97, 38, 15, 00, 40, 00, 75, 04, 05, 07, 78, 52, 12, 50, 77, 91, 8},
		[20]int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 00},
		[20]int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65},
		[20]int{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91},
		[20]int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		[20]int{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		[20]int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		[20]int{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		[20]int{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		[20]int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92},
		[20]int{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95},
		[20]int{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57},
		[20]int{86, 56, 00, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		[20]int{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40},
		[20]int{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		[20]int{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		[20]int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16},
		[20]int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		[20]int{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54},
		[20]int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48}}
	biggestProd := 0

	for row := 0; row < 20; row++ {
		for col := 0; col < 20; col++ {

			// going right
			prod := 1
			for r := 0; r < 4 && row+r < 20; r++ {
				prod *= matrix[row+r][col]
			}
			if prod > biggestProd {
				biggestProd = prod
			}

			// going down
			prod = 1
			for d := 0; d < 4 && col+d < 20; d++ {
				prod *= matrix[row][col+d]
			}
			if prod > biggestProd {
				biggestProd = prod
			}

			// going down diagonal
			prod = 1
			for d := 0; d < 4 && col+d < 20 && row+d < 20; d++ {
				prod *= matrix[row+d][col+d]
			}
			if prod > biggestProd {
				biggestProd = prod
			}

			// going up diagonal
			prod = 1
			for d := 0; d < 4 && col-d >= 0 && row+d < 20; d++ {
				prod *= matrix[row+d][col-d]
			}
			if prod > biggestProd {
				biggestProd = prod
			}
		}
	}

	return biggestProd
}

func problem12() int {
	triplet := 0
	dividers := 0
	for i := 1; ; i++ {
		triplet += i

		dividers = 1
		pFactors := PrimeFactors(uint64(triplet))
		pFactors = append(pFactors, 1)
		lastPrime := 0
		prod := 1

		for _, p := range pFactors {
			if lastPrime == int(p) {
				prod++
			} else {
				dividers *= prod
				prod = 2
				lastPrime = int(p)
			}
		}

		if dividers >= 500 {
			break
		}
	}

	return triplet
}

func problem13() int {
	lines := ReadLinesFrom("src/github.com/reicher/euler/data/prob13.txt")
	Sum := big.NewInt(0)
	for _, n := range lines {
		bigNum := new(big.Int)
		bigNum.SetString(n, 10)
		Sum = Sum.Add(bigNum, Sum)
	}
	SumString := Sum.String()
	result, _ := strconv.Atoi(SumString[0:10])
	return result
}

// Longest Collatz sequence
func problem14() int {
	longest := 0
	startOfLongest := 0

	for i := 1; i < 1000000; i++ {
		sequenceLength := GetCollatzSequenceLength(i)
		if longest < sequenceLength {
			longest = sequenceLength
			startOfLongest = i
		}
	}

	return startOfLongest
}

var Solutions map[int]func() int

func init() {
	Solutions = make(map[int]func() int)
	Solutions[1] = problem1
	Solutions[2] = problem2
	Solutions[3] = problem3
	Solutions[4] = problem4
	Solutions[5] = problem5
	Solutions[6] = problem6
	Solutions[7] = problem7
	Solutions[8] = problem8
	Solutions[9] = problem9
	Solutions[10] = problem10
	Solutions[11] = problem11
	Solutions[12] = problem12
	Solutions[13] = problem13
	Solutions[14] = problem14
}

func Solve(n int) int {
	start := time.Now()

	answer := Solutions[n]()

	elapsed := time.Since(start)
	fmt.Printf("\nProblem %d solved in %s : %d\n", n, elapsed, answer)
	return answer
}
