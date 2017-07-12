package factorial_test

import (
	"math/big"
	"testing"
	"runtime"

	"github.com/bhansconnect/factorial"
)

var fact100, _ = new(big.Int).SetString("93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000", 10)

var tests = []struct {
	n        int64
	expected *big.Int
}{
	{-1, big.NewInt(-1)},
	{0, big.NewInt(1)},
	{1, big.NewInt(1)},
	{5, big.NewInt(120)},
	{10, big.NewInt(3628800)},
	{100, fact100},
}

func TestParallelFactorial(t *testing.T) {

	runtime.GOMAXPROCS(2)
	for _, tt := range tests {
		actual := factorial.ParallelFactorial(tt.n)
		if actual.String() != tt.expected.String() {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
	runtime.GOMAXPROCS(3)
	for _, tt := range tests {
		actual := factorial.ParallelFactorial(tt.n)
		if actual.String() != tt.expected.String() {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkParallelFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.ParallelFactorial(10000)
	}
}

func TestFactorial(t *testing.T) {
	for _, tt := range tests {
		actual := factorial.Factorial(tt.n)
		if actual.String() != tt.expected.String() {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial.Factorial(10000)
	}
}
