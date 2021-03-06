package main

import "os"

func main() {
	src, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	bs := make([]byte, 9)
	src.Read(bs)
	dst.Write(bs)

	src.Close()

	os.Remove("text.txt")
}
