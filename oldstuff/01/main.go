package main

import (
	"fmt"
	"go/build"
)

func main() {
	fmt.Println(build.Default.GOOS)
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
