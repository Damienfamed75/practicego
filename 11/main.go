package main

import (
	"fmt"
)

// fizzbuzz challenge
//====================
// For every multiple of 3 print 'fizz'
// For every multiple of 5 print 'buzz'
// If multiple by both then print 'fizzbuzz'
// If multiple by neither then do not print.
// Do this up to 100 numbers...

func main() {
	i := 0
	for {
		i++
		if i > 100 {
			break
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "fizz")
			continue

		}
		if i%5 == 0 {
			fmt.Println(i, "buzz")
			continue
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
	}
}
