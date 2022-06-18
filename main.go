package main

import (
	"fmt"
	"zoologic/constants"
	"zoologic/src"
)

func main() {
	first := src.GetSpeciesById(constants.BearsId)
	fmt.Println(first)
}
