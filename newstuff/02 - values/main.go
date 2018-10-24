package main

// --------------------------------------------
// This example shows different ways to interact
// with different types in Go. This includes
// addition with strings, addition with ints
// and division with floats. Another thing this
// example does is showcase the boolean operators
// such as &&, ||, and !.
// --------------------------------------------

import "fmt"

func main() {
	// strings can be 'concatenated' with +
	fmt.Println("go" + "lang")

	// Integers
	fmt.Println("1+1 =", 1+1)

	// Floats
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans with boolean operators.

	// "return true if (true and false are true)"
	// '&&' = and
	fmt.Println(true && false)

	// "return true if (true or false are true)"
	// '||' = or
	fmt.Println(true || false)

	// "return true if (not true is true)"
	// '!' = not
	fmt.Println(!true)
}
