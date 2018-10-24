package main

import (
	"fmt"
)

func getNames() map[string]string {
	return map[string]string{
		"mason":  "grimes",
		"peter":  "jelen",
		"noah":   "mastaglio",
		"sierra": "winter",
		"thom":   "patterson",
		"eddie":  "estevez",
	} // second way
}

func main() {
	// first way
	names := make(map[string]string)
	names["damien"] = "stamates"
	names["kyle"] = "clavette"
	names["john"] = "knecht"
	names["jeff"] = "kuang"
	names["bahram"] = "samani"
	names["richard"] = "lowery"

	internNames := getNames()

	for fn, ln := range names {
		fmt.Println(fn + " " + ln)
	}

	fmt.Println()
	if val, exists := internNames["noah"]; exists {
		delete(internNames, "noah") // Do you know Noah? No you don't.
		fmt.Println("deleted " + "noah " + val)
	}

	fmt.Println()
	for fn, ln := range internNames {
		fmt.Println(fn + " " + ln)
	}
}
