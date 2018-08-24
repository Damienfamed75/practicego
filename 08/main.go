package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your name? ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("Have a Nice Day!", text)
}
