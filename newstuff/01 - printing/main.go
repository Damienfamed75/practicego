// package main means the current package you
// are working in is the 'main' package. This
// also means that this file requires a 'main'
// function inside of it to run.
package main

// --------------------------------------------
// This example shows off some basic printing
// using Golang's internal "fmt" package.
// --------------------------------------------

// When you're going to be using functions
// from outside of the current package then you
// need to 'import' them. This is usually done
// automatically by VSCode/Any other IDE.
import "fmt"

// main will print out some example sentences
// using the three Printing functions inside
// the "fmt" package.
func main() {

	// Print doesn't create a new line after
	// being called. In order to make a new line
	// we use the escape character '\n'.
	fmt.Print("Hello, ")
	fmt.Print("World!\n")

	// Println creates a new lines after
	// being called so you don't have to use
	// the escape character for a new line.
	fmt.Println("Hello, ") // Output: Hello, \n
	fmt.Println("World!")  // Output: World!\n

	// age will represent the age of myself.
	// This variable is defaulted type 'int'.
	// You can see this by hovering your mouse
	// over the variable name.
	age := 19

	// Printf is a formatting alternate to
	// Print. This is an example of using an
	// escape character '%v' which simply means,
	// "to put a 'value' here." That value is
	// then placed on the right hand side.
	fmt.Printf("I am %v years old", age)
}
