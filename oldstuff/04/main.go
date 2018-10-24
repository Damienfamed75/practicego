package main

import "fmt"

func makeSummer() func(sm ...int) int {
	return func(sm ...int) int {
		var total int
		for _, nm := range sm {
			total += nm
		}
		return total
	}
}

func main() {
	summer := makeSummer()
	fmt.Println(summer(2, 4, 6, 4, 4, 6, 2, 7, 9, 23, 84))
}
