package src

import (
	"fmt"
	"strings"
	"zoologic/data"
	"zoologic/types"
)

func GetAnimalsTotalQuantities() []types.AnimalCounterOutput {
	var animalsTotalQuantities []types.AnimalCounterOutput

	for _, animal := range data.Zoo.Species {
		animalData := types.AnimalCounterOutput{Specie: animal.Name, Quantity: len(animal.Residents)}
		animalsTotalQuantities = append(animalsTotalQuantities, animalData)
	}

	return animalsTotalQuantities
}

func CountAnimals(animal *types.AnimalCounterInput) []types.AnimalCounterOutput {
	var animalsCounter []types.AnimalCounterOutput

	if animal == nil {
		allAnimals := GetAnimalsTotalQuantities()
		animalsCounter = append(animalsCounter, allAnimals...)
		return animalsCounter
	}

	if animal.Sex == nil {
		fmt.Println("here")
		for _, specie := range data.Zoo.Species {
			temp := *animal.Specie
			if strings.EqualFold(specie.Name, temp) {
				animalTotalQuantity := types.AnimalCounterOutput{Specie: specie.Name, Quantity: len(specie.Residents)}
				animalsCounter = append(animalsCounter, animalTotalQuantity)
			}
		}
		return animalsCounter
	}

	for _, specie := range data.Zoo.Species {
		counter := 0
		tempSpecie, tempSex := *animal.Specie, *animal.Sex
		if strings.EqualFold(specie.Name, tempSpecie) {
			for _, resident := range specie.Residents {
				if strings.EqualFold(resident.Sex, tempSex) {
					counter++
				}
			}
			animalTotalQuantity := types.AnimalCounterOutput{Specie: specie.Name, Quantity: counter}
			animalsCounter = append(animalsCounter, animalTotalQuantity)
		}
	}

	return animalsCounter
}
