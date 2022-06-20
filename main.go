package main

import (
	"fmt"
	"zoologic/src"
	"zoologic/types"
)

func main() {
	// first := src.GetSpeciesById(constants.BearsId)
	// fmt.Println(first)

	// second := src.GetAnimalsOlderThan("Bears", 2)
	// fmt.Println(second)

	// third := src.GetEmployeeByName("burl")
	// fmt.Println(third)

	// fourth, err := src.GetRelatedEmployees(constants.StephanieId)
	// fmt.Println(fourth, err)

	name := "giraffes"
	sex := "male"
	animal := types.AnimalCounterInput{Specie: &name, Sex: &sex}
	fifth := src.CountAnimals(&animal)
	fmt.Println(fifth)
}
