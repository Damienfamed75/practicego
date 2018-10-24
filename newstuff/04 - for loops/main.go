package main

import "fmt"

func main() {

	// declare our new variable 'i'
	// which equals 1.
	i := 1

	// run this loop while
	// 'i' is less than or equal to 3.
	for i <= 3 {
		fmt.Println("i =", i)

		// another way of doing this is:
		// 'i++' or 'i += 1'
		i = i + 1
	}

	// j = 7.
	// run while j is less than or equal to 9.
	// after every loop add 1 to j using 'j++'
	for j := 7; j <= 9; j++ {
		fmt.Println("j =", j)
	}

	// no condition given to 'for'
	// this means it will forever run
	// until it is broken.
	for {
		fmt.Println("infinite loop until break")
		break
	}

	// n = 0
	// run while n is less than or equal to 5.
	// after every loop add 1 to n using 'n++'
	for n := 0; n <= 5; n++ {
		// if n is an even number
		// then 'continue' to the
		// next loop.
		if n%2 == 0 {
			continue
		}
		fmt.Println("n =", n)
	}
}
