package src

import (
	"sort"
	"zoologic/data"
	"zoologic/types"
)

func GetAnimalsNames() map[string][]string {
	namedAnimals := make(map[string][]string)

	for _, specie := range data.Zoo.Species {
		for _, resident := range specie.Residents {
			namedAnimals[specie.Name] = append(namedAnimals[specie.Name], resident.Name)
		}
	}

	return namedAnimals
}

func GetAnimalsStructs(animalSpecie string, shouldSort bool) types.Animal {
	specie := animalSpecie
	residents := GetAnimalsNames()[animalSpecie]

	if shouldSort {
		sort.Strings(residents)
	}

	structure := types.Animal{Specie: specie, Residents: residents}

	return structure
}

func GetAnimalsLocalization() types.MappedAnimals {
	mappedAnimals := types.MappedAnimals{}

	for _, specie := range data.Zoo.Species {
		animalData := types.Animal{}
		if specie.Location == "NE" {
			animalData.Specie = specie.Name
			mappedAnimals.NE = append(mappedAnimals.NE, animalData)
		}
		if specie.Location == "NW" {
			animalData.Specie = specie.Name
			mappedAnimals.NW = append(mappedAnimals.NW, animalData)
		}
		if specie.Location == "SE" {
			animalData.Specie = specie.Name
			mappedAnimals.SE = append(mappedAnimals.SE, animalData)
		}
		if specie.Location == "SW" {
			animalData.Specie = specie.Name
			mappedAnimals.SW = append(mappedAnimals.SW, animalData)
		}
	}

	return mappedAnimals
}

func GetAnimalCompleteMap(shouldSort bool) types.MappedAnimals {
	mappedAnimals := types.MappedAnimals{}

	lionsStruct := GetAnimalsStructs("lions", shouldSort)
	mappedAnimals.NE = append(mappedAnimals.NE, lionsStruct)

	giraffesStruct := GetAnimalsStructs("giraffes", shouldSort)
	mappedAnimals.NE = append(mappedAnimals.NE, giraffesStruct)

	tigersStruct := GetAnimalsStructs("tigers", shouldSort)
	mappedAnimals.NW = append(mappedAnimals.NW, tigersStruct)

	bearsStruct := GetAnimalsStructs("bears", shouldSort)
	mappedAnimals.NW = append(mappedAnimals.NW, bearsStruct)

	elephantsStruct := GetAnimalsStructs("elephants", shouldSort)
	mappedAnimals.NW = append(mappedAnimals.NW, elephantsStruct)

	penguinsStruct := GetAnimalsStructs("penguins", shouldSort)
	mappedAnimals.SE = append(mappedAnimals.SE, penguinsStruct)

	ottersStruct := GetAnimalsStructs("otters", shouldSort)
	mappedAnimals.SE = append(mappedAnimals.SE, ottersStruct)

	frogsStruct := GetAnimalsStructs("frogs", shouldSort)
	mappedAnimals.SW = append(mappedAnimals.SE, frogsStruct)

	snakesStruct := GetAnimalsStructs("snakes", shouldSort)
	mappedAnimals.SW = append(mappedAnimals.SE, snakesStruct)

	return mappedAnimals
}

func GetAnimalsMap(options *types.GetAnimalsMapInput) types.MappedAnimals {
	if options == nil {
		mappedAnimals := GetAnimalsLocalization()
		return mappedAnimals
	}

	if options.IncludeNames == nil && options.Sorted == nil {
		mappedAnimals := GetAnimalsLocalization()
		return mappedAnimals
	}

	if options.IncludeNames == nil {
		mappedAnimals := GetAnimalsLocalization()
		return mappedAnimals
	}

	if *options.IncludeNames && *options.Sorted {
		shouldSort := true
		mappedAnimals := GetAnimalCompleteMap(shouldSort)

		return mappedAnimals
	}

	shouldSort := false
	mappedAnimals := GetAnimalCompleteMap(shouldSort)

	return mappedAnimals
}
