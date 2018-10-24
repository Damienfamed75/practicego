package objects

import "fmt"

type ObjectOne struct {
	FName string
	LName string
}

func NewObjectOne(fName string, lName string) *ObjectOne {
	objOne := &ObjectOne{
		FName: fName,
		LName: lName,
	}

	return objOne
}

func (o *ObjectOne) UpdateNames(fName string, lName string) {
	o.FName = fName
	o.LName = lName
}

func (o *ObjectOne) Update() {
	// fmt.Println("HI")
	fmt.Println("Hi I'm", o.FName, o.LName)
}

func (o *ObjectOne) New(fName string, lName string) *ObjectOne {
	o.FName = fName
	o.LName = lName

	return o
}
