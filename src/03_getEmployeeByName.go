package src

import (
	"strings"
	"zoologic/data"
	"zoologic/types"
)

func GetEmployeeByName(employeeName string) types.Employee {
	var targetEmployee types.Employee

	for _, employee := range data.Zoo.Employees {
		if strings.EqualFold(employee.FirstName, employeeName) || strings.EqualFold(employee.LastName, employeeName) {
			targetEmployee = employee
			break
		}
	}

	return targetEmployee
}
