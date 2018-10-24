package main

import (
	"fmt"

	"github.com/damienfamed75/practicego/00/shapes"
)

func main() {
	r := &shapes.Rect{Width: 10, Height: 5}

	fmt.Println("area: ", r.Area())
	fmt.Println("perim:", r.Perim())
}
