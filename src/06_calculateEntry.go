package src

import (
	"math"
	"zoologic/data"
	"zoologic/types"
)

func CountEntrants(entrantsList []types.EntrantInput) types.EntrantOutput {
	var mappedEntrants types.EntrantOutput

	for _, entrant := range entrantsList {
		if entrant.Age < 18 {
			mappedEntrants.Child++
		} else if entrant.Age >= 50 {
			mappedEntrants.Senior++
		} else {
			mappedEntrants.Adult++
		}
	}

	return mappedEntrants
}

func CalculateTotalPrice(entrantsList []types.EntrantInput) float64 {
	prices := data.Zoo.Prices

	entrants := CountEntrants(entrantsList)

	totalPrice := prices.Adult*float64(entrants.Adult) + prices.Senior*float64(entrants.Senior) + prices.Child*float64(entrants.Child)

	return math.Round(totalPrice*100) / 100
}
