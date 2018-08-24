package main

import (
	"fmt"
)

// for

// if

// continue

// break

func main() {
	i := 0
	for {
		i++
		if i%2 != 0 {
			continue

		}
		fmt.Println(i)
		if i > 50 {
			break

		}
	}
}
