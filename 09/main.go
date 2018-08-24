package main

import (
	"github.com/Damienfamed75/practicego/09/examiner"
)

func main() {
	e := examiner.NewExaminer()

	if !e.HasExamined() {
		e.ExamineAndTell("red", false)
	}

	e.ExamineAndTell("yellow", true)

	e.ExamineAndTell("green", false)

	e.ExamineAndTell("orange", true)
}
