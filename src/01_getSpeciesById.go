package src

import (
	"zoologic/data"
	"zoologic/types"
)

func GetSpeciesById(ids ...string) []types.ZooSpecie {

	foundSpecies := make([]types.ZooSpecie, 0, len(ids))

	for _, paramId := range ids {

		for _, specie := range data.Zoo.Species {

			if specie.Id == paramId {

				foundSpecies = append(foundSpecies, specie)

			}

		}

	}

	return foundSpecies
}
