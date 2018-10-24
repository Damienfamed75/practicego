package main

import "fmt"

func main() {
	// declare an empty array of integers
	// that is 5 elements in length.
	var a [5]int
	fmt.Println("emp:", a)

	// the fifth element is set to 100.
	// if you want to get the last element
	// of an array you use: a[len(a)-1]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len() is a built-in Golang function.
	// it's used to get the length of
	// arrays, slices, and maps.
	fmt.Println("len:", len(a))

	// declare an array with hardcoded values.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// create a two dimensional array.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
