package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Damienfamed75/practicego/10/calculator"
)

func main() {
	calc := calculator.New()

	// get user input
	// store first variable
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Whats The First Number!")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print("Second Number!")
	texto, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q\n", text)
	fmt.Printf("%q\n\n", texto)

	text = strings.TrimSuffix(text, "\r\n")
	texto = strings.TrimSuffix(texto, "\r\n")

	numOne, _ := strconv.Atoi(text)
	numTwo, _ := strconv.Atoi(texto)

	fmt.Println(calc.Add(numOne, numTwo))
	// get user input
	// store second variable

	// print added numbers

	// fmt.Println(calc.Add(1, 4))
}
