package objects

import "fmt"

type ObjectTwo struct {
	FName string
	MName string
	LName string
}

func NewObjectTwo(fName string, mName string, lName string) *ObjectTwo {
	objTwo := &ObjectTwo{
		FName: fName,
		MName: mName,
		LName: lName,
	}

	return objTwo
}

func (o *ObjectTwo) Update() {
	// fmt.Println("HELLO")
	fmt.Println("Hello I'm", o.FName, o.MName, o.LName)
}
