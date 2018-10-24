package main

import (
	"fmt"
	"math"
)

// const declares a constant value.

// a const statement can appear anywhere a var
// statement can.

// s is now a constant across this entire package.
// The package being "main"
const s string = "constant"

func main() {
	// print the global constant 's'
	fmt.Println(s)

	// call a function that will print 's' also.
	printS()

	const n = 500000000

	// Constant expressions perform arithmetic
	// with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(500000000))
}

// printS is a function that will print
// the 's' constant which is a package global.
func printS() {
	fmt.Println(s)
}
