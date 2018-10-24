package main

import "fmt"

func main() {
	av := average(20, 10, 40, 46, 94, 42)
	fmt.Println(av)
}

func average(sf ...float64) float64 {
	var total float64
	for _, num := range sf {
		total += num
	}
	return total / float64(len(sf))
}
