package main

import (
	"fmt"
)

func acc(c chan int, coords []int) {
	for _, v := range coords {
		c <- v
	}
}

func main() {
	c := make(chan int)
	coords := []int{1, 2, 7, 3, 1, 6, 7}

	for range coords {
		fmt.Println(<-c)
	}
}
