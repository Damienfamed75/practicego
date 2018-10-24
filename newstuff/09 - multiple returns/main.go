package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// in Golang you can't have variables
	// that are left unused. So if you have a
	// function that returns multiple things and you
	// only want one value then you use _ to omit
	// a value. For instance I only want the second
	// value from this function. So I omit the first
	// return value with an _.
	_, c := vals()
	fmt.Println(c)
}
