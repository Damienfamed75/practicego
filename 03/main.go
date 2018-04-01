package main

import "fmt"

func main() {
	data := []int{12, 4, 2, 3, 5, 8, 2, 6}
	num := sum(data...)
	fmt.Println(num)
}

func sum(nf ...int) int {
	var total int
	for _, nm := range nf {
		total += nm
	}
	return total
}
