package main

import (
	"fmt"
	"time"
)

func main() {
	name := "damien"

	switch name {
	case "jacob":
		fmt.Println("You're not programming so get away.")
	case "nathan":
		fmt.Println("Hey Nathan how are you doing with all this Golang?")
	case "damien":
		fmt.Println("Great... the dunce is here to read his own code.")
	}

	// ==================================================================

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday...")
	}

	// ==================================================================

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// interface{} is used as a generic value.
// any kind of variable can be passed if it's
// an interface.
func whatAmI(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}
