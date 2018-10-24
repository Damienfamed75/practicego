package main

import "fmt"

func numbers() {
	defer fmt.Println(2)

	fmt.Println(1)
} // end of the scope

// defer means that the function will be called at the end of the scope.
func main() {
	defer fmt.Println("Goodbye!")

	numbers()

	fmt.Println("Hello!")

} // end of the scope
