package main

import (
	"fmt"
	"zoologic/constants"
	"zoologic/src"
)

func main() {
	// first := src.GetSpeciesById(constants.BearsId)
	// fmt.Println(first)

	// second := src.GetAnimalsOlderThan("Bears", 2)
	// fmt.Println(second)

	// third := src.GetEmployeeByName("burl")
	// fmt.Println(third)

	fourth, err := src.GetRelatedEmployees(constants.StephanieId)
	fmt.Println(fourth, err)
}
