package src

import (
	"errors"
	"zoologic/data"
)

func isManager(personId string) bool {
	isManager := false

	for _, employee := range data.Zoo.Employees {
		for _, manager := range employee.Managers {
			if manager == personId {
				isManager = true
				break
			}
		}
	}

	return isManager
}

func GetRelatedEmployees(personId string) (team []string, noManagerError error) {
	if manager := isManager(personId); !manager {
		return nil, errors.New("o id inserido não é de uma pessoa colaboradora gerente")
	}

	for _, employee := range data.Zoo.Employees {
		for _, manager := range employee.Managers {
			if manager == personId {
				completeName := employee.FirstName + " " + employee.LastName
				team = append(team, completeName)
			}
		}
	}

	return team, nil
}
