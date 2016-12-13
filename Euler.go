package euler

import "math"
import (
	"fmt"
	"time"
	"strconv"
)

func problem1() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if(math.Mod(float64(i), 3) == 0 || math.Mod(float64(i), 5) == 0){
			sum += i;
		}
	}
	return sum;
}

func problem2() int {
	fibs := GetFibonacci(4000000);
	sum := 0;
	for _, f := range fibs {
		if f%2 == 0 {
			sum += f;
		}
	}
	return sum;
}

func problem3() int {
	 n := uint64(600851475143)
	factors := PrimeFactors(n);
	return int(factors[len(factors)-1]);
}

func problem4() int {
	for i := 999*999; i >= 0; i-- {
		if IsPalindrome(i) {
			for j := 999; j > 99; j-- {
				if i%j == 0 && i/j <= 999 {
					return i;
				}
			}
		}
	}
	return 0;
}

func problem5() int {
	var factors [20]int
	var maxFactors [20]int

	for i:= 0; i < len(maxFactors); i++ {
		maxFactors[i] = 0;
	}

	for i := 2; i < len(factors); i++ {
		for i:= 0; i < len(maxFactors); i++ {
			factors[i] = 0;
		}

		for _, f := range PrimeFactors(uint64(i)) {
			factors[f]++;
		}

		for j, f := range factors {
			if f > maxFactors[j] {
				maxFactors[j] = f;
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
	n := 100;

	sumOfSquare := 0;
	SquareOfSum := 0;

	for i := 1; i <= n; i++ {
		sumOfSquare += i*i;
		SquareOfSum += i;
	}
	SquareOfSum = SquareOfSum * SquareOfSum;

	return SquareOfSum - sumOfSquare
}

func problem7() int {
	loadsOfPrimes := GetFirstPrimes(10001)
	return loadsOfPrimes[len(loadsOfPrimes)-1];
}

func problem8() int {
	number := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	maxProduct := 0;
	cl := 13;

	for i := 0; i <= len(number)-cl; i++ {
		product, _ := strconv.Atoi( string(number[i]) )
		for j := i+1; j < cl+i || j > len(number); j++ {
			num, _ := strconv.Atoi( string(number[j]) )
			product *= num;
		}


		if product > maxProduct {
			maxProduct = product;
		}
		if (i + cl) >= len(number) {
			break;
		}
	}

	return maxProduct;
}

var Solutions map[int] func() int

func init () {
	Solutions = make(map[int] func() int)
	Solutions[1] = problem1
	Solutions[2] = problem2
	Solutions[3] = problem3
	Solutions[4] = problem4
	Solutions[5] = problem5
	Solutions[6] = problem6
	Solutions[7] = problem7
	Solutions[8] = problem8
}

func Solve(n int) int{
	start := time.Now()

	answer := Solutions[n]();

	elapsed := time.Since(start)
	fmt.Printf("\nProblem %d solved in %s : %d\n", n, elapsed, answer)
	return answer;
}
