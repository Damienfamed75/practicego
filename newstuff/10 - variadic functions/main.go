package main

import "fmt"

// Here's a function that will take
// an arbitrary number of ints as
// arguments.

// when inside of the function
// the argument 'nums' acts like
// a slice/array. So using len()
// and for loops work.
func sum(nums ...int) {
	// print the slice/array of ints
	fmt.Print(nums, " ")

	total := 0

	// _ is the index
	// for every number in the
	// array of nums.
	for _, num := range nums {
		// total = total + num
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	// using ... after an array or slice it will
	// lay out the array like:
	// sum(1, 2, 3, 4)
	sum(nums...)
}
