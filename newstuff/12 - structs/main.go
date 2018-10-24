package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	// You don't need to include every single field
	// of the struct in order to make a struct.
	fmt.Println(&person{name: "Fred"})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// using pointers we
	// can directly affect
	// 's' and its fields
	sp := &s
	fmt.Println("&Sean age =", sp.age)

	sp.age = 51
	fmt.Println("&Sean age =", sp.age)
	fmt.Println("Sean age =", s.age)
}
