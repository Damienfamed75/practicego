package examiner

import (
	"fmt"

	"github.com/Damienfamed75/practicego/09/apple"
)

// Examiner examines to see what the apple is
type Examiner struct {
	Apple       *apple.Apple
	hasExamined bool // bool (booleans) contain a true or false value
}

// NewExaminer returns a new Examiner to examine apple
func NewExaminer() *Examiner {

	return &Examiner{
		hasExamined: false,
	}
}

// Examine the apple inside of the Examiner object
func (e *Examiner) Examine(col string, worm bool) {
	e.hasExamined = true
	e.Apple = apple.NewApple(col, worm)
}

// Stupid prints out the stupid apple's color
func (e *Examiner) Stupid() {
	fmt.Println(e.Apple.Color)
}

// HasExamined gets the value of 'hasExamined'
func (e *Examiner) HasExamined() bool {
	return e.hasExamined
}

// Worm prints out if its true or false
func (e *Examiner) Worm() {
	fmt.Println(e.Apple.Worm())
}

// ExamineAndTell examines a new apple and prints out a description of it.
func (e *Examiner) ExamineAndTell(col string, worm bool) {
	e.Examine(col, worm)
	fmt.Println(e.Apple.Color, e.Apple.Worm())
}
