package main

import (
	"fmt"

	"github.com/Damienfamed75/practicego/02/something"
)

func main() {
	added := add(1, 2, 4, 5, 5, 12)

	something.SayMyName()
	something.SayStupid()

	fmt.Println(added)
}

func add(sf ...int) int {
	total := 0
	for _, num := range sf {

		total = total + num
	}
	return total
}
