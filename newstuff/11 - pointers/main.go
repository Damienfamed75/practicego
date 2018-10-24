package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// This example showcases how to use
// pointers and what a pointer is.

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// When using & you grab the
	// memory address of the variable.
	// This address is an address like a house address.
	// It tells the computer where to find this variable
	// in the RAM.
	fmt.Println("pointer:", &i)
}
