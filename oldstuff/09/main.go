package main

import (
	"github.com/damienfamed75/practicego/09/objects"
	"github.com/damienfamed75/practicego/09/types"
)

var (
	u types.Updater
)

func main() {
	personOne := objects.NewObjectOne("Damien", "Stamates")
	personTwo := objects.NewObjectTwo("Nathan", "Daniel", "Melsheimer")

	updaters := []types.Updater{personOne, personTwo}

	update(updaters)

	personOne.FName = "Someone"
	personOne.LName = "else"

	update(updaters)
}

func update(updaters []types.Updater) {
	for _, u := range updaters {
		u.Update()
	}
}
