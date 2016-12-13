package euler_test

import (
	"testing"
	"github.com/reicher/euler"
)

var answers = []int {233168, 4613732, 6857, 906609,232792560, 25164150, 104743, 23514624000}

func TestFactorization(t *testing.T) {
	n := 59855985
	factors := euler.PrimeFactors(uint64(n))

	product := 1
	for i := 0; i < len(factors); i++ {
		product *= int(factors[i]);
	}

	if n != product {
		t.Errorf("Prime factors of %d is not %v", n,  factors)
	}
}

func TestProblems(t *testing.T) {
	if len(answers) != len(euler.Solutions) {
		t.Errorf("Missing solutions!");
		return;
	}

	for p := 1; p <= len(euler.Solutions); p++ {
		actual := euler.Solve(p);
		target := answers[p-1];
		if actual != target {
			t.Errorf("Problem %d\nTarget: %d\nActual: %d", p, target, actual);
		}
	}
}