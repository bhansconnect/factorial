package factorial

import (
	"math/big"
	"runtime"
)

//ParallelFactorial calculates the factorial of n
//The calculation is done concurrently
//Creates a go routine for each proc
//If n < 0, -1 is returned
func ParallelFactorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n < 0 {
		return big.NewInt(-1)
	}
	procs := int64(runtime.GOMAXPROCS(0))
	if n < procs {
		return Factorial(n)
	}

	outRange := make(chan *big.Int, procs)

	for i := int64(0); i < procs; i++ {
		go multrange(i*n/procs+1, (i+1)*n/procs, outRange)
	}

	in := outRange
	for procs > 1 {
		out := make(chan *big.Int, procs/2+1)
		odd := false
		if procs%2 == 1 {
			odd = true
		}
		procs /= 2
		for i := int64(0); i < procs; i++ {
			go mult(in, out)
		}
		if odd {
			out <- <-in
			procs++
		}
		in = out
	}
	return <-in
}

func mult(in <-chan *big.Int, out chan<- *big.Int) {
	total := <-in
	n := <-in
	total.Mul(total, n)
	out <- total
}

func multrange(x int64, y int64, out chan<- *big.Int) {
	total := big.NewInt(0)
	total.MulRange(x, y)
	out <- total
}

//Factorial calculates the factorial of n
//The calculation is done sequentially
//If n < 0, -1 is returned
func Factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n < 0 {
		return big.NewInt(-1)
	}
	out := big.NewInt(1)
	out.MulRange(1, n)
	return out
}
