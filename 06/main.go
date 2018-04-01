package main

import "fmt"

func main() {
	data := []float64{9, 22, 34, 5, 2, 54, 10, 14, 18, 20, 32}

	fmt.Println(average(data[0:5]...)) // accessing 0-5 in the slice. Not including the 5th element
	fmt.Println(average(data...))
}

func average(sf ...float64) float64 {
	var total float64
	for _, nm := range sf {
		total += nm
	}
	return total / float64(len(sf))
}
