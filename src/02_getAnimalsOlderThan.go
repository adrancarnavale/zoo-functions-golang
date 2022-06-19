package src

import (
	"strings"
	"zoologic/data"
	"zoologic/types"
)

func GetAnimalsOlderThan(specieName string, minimalAge int) bool {
	areAllOlderThan := true
	var foundSpecie types.ZooSpecie

	for _, specie := range data.Zoo.Species {
		if strings.EqualFold(specie.Name, specieName) {
			foundSpecie = specie
			break
		}
	}

	if foundSpecie.Id == "" {
		areAllOlderThan = false
	}

	for _, resident := range foundSpecie.Residents {
		if resident.Age < minimalAge {
			areAllOlderThan = false
			break
		}
	}

	return areAllOlderThan

}
