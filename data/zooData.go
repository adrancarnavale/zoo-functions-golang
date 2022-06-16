package data

import (
	"zoologic/constants"
	"zoologic/types"
)

const (
	lionId      = constants.LionId
	ottersId    = constants.OttersId
	elephantsId = constants.ElephantsId
	snakesId    = constants.SnakesId
	frogsId     = constants.FrogsId
	bearsId     = constants.BearsId
	tigersId    = constants.TigersId
	penguinsId  = constants.PenguinsId
	giraffesId  = constants.GiraffesId

	stephanieId = constants.StephanieId
	burlId      = constants.BurlId
	olaId       = constants.OlaId
	nigelId     = constants.NigelId
	wilburnId   = constants.WilburnId
	sharondaId  = constants.SharondaId
	ardithId    = constants.ArdithId
	emeryId     = constants.EmeryId
)

type (
	Zoologic  = types.Zoologic
	Day       = types.Day
	Employee  = types.Employee
	ZooSpecie = types.ZooSpecie
	Resident  = types.Resident
	Prices    = types.Prices
	Hours     = types.Hours
)

var Zoo = Zoologic{
	Species: []ZooSpecie{
		{
			Id:           lionId,
			Name:         "lions",
			Popularity:   4,
			Location:     "NE",
			Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
			Residents: []Resident{
				{
					Name: "Zena",
					Sex:  "female",
					Age:  12,
				},
				{
					Name: "Maxwell",
					Sex:  "male",
					Age:  15,
				},
				{
					Name: "Faustino",
					Sex:  "male",
					Age:  7,
				},
				{
					Name: "Dee",
					Sex:  "female",
					Age:  14,
				},
			},
		},
		{
			Id:           tigersId,
			Name:         "tigers",
			Popularity:   5,
			Location:     "NW",
			Availability: []string{"Wednesday", "Friday", "Saturday", "Tuesday"},
			Residents: []Resident{
				{
					Name: "Shu",
					Sex:  "female",
					Age:  19,
				},
				{
					Name: "Esther",
					Sex:  "female",
					Age:  17,
				},
			},
		},
		{
			Id:           bearsId,
			Name:         "bears",
			Popularity:   5,
			Location:     "NW",
			Availability: []string{"Tuesday", "Wednesday", "Sunday", "Saturday"},
			Residents: []Resident{
				{
					Name: "Hiram",
					Sex:  "male",
					Age:  4,
				},
				{
					Name: "Edwardo",
					Sex:  "male",
					Age:  4,
				},
				{
					Name: "Milan",
					Sex:  "male",
					Age:  4,
				},
			},
		},
		{
			Id:           penguinsId,
			Name:         "penguins",
			Popularity:   4,
			Location:     "SE",
			Availability: []string{"Tuesday", "Wednesday", "Sunday", "Saturday"},
			Residents: []Resident{
				{
					Name: "Joe",
					Sex:  "male",
					Age:  10,
				},
				{
					Name: "Tad",
					Sex:  "male",
					Age:  12,
				},
				{
					Name: "Keri",
					Sex:  "female",
					Age:  2,
				},
				{
					Name: "Nicholas",
					Sex:  "male",
					Age:  2,
				},
			},
		},
		{
			Id:           ottersId,
			Name:         "otters",
			Popularity:   4,
			Location:     "SE",
			Availability: []string{"Friday", "Thursday", "Wednesday", "Saturday"},
			Residents: []Resident{
				{
					Name: "Neville",
					Sex:  "male",
					Age:  9,
				},
				{
					Name: "Lloyd",
					Sex:  "male",
					Age:  8,
				},
				{
					Name: "Mercedes",
					Sex:  "female",
					Age:  9,
				},
				{
					Name: "Margherita",
					Sex:  "female",
					Age:  10,
				},
			},
		},
		{
			Id:           frogsId,
			Name:         "frogs",
			Popularity:   2,
			Location:     "SW",
			Availability: []string{"Saturday", "Friday", "Thursday", "Wednesday"},
			Residents: []Resident{
				{
					Name: "Cathey",
					Sex:  "female",
					Age:  3,
				},
				{
					Name: "Annice",
					Sex:  "female",
					Age:  2,
				},
			},
		},
		{
			Id:           snakesId,
			Name:         "snakes",
			Popularity:   3,
			Location:     "SW",
			Availability: []string{"Sunday", "Saturday", "Friday", "Thursday"},
			Residents: []Resident{
				{
					Name: "Paulette",
					Sex:  "female",
					Age:  5,
				},
				{
					Name: "Bill",
					Sex:  "male",
					Age:  6,
				},
			},
		},
		{
			Id:           elephantsId,
			Name:         "elephants",
			Popularity:   5,
			Location:     "NW",
			Availability: []string{"Friday", "Saturday", "Sunday", "Tuesday"},
			Residents: []Resident{
				{
					Name: "Ilana",
					Sex:  "female",
					Age:  11,
				},
				{
					Name: "Orval",
					Sex:  "male",
					Age:  15,
				},
				{
					Name: "Bea",
					Sex:  "female",
					Age:  12,
				},
				{
					Name: "Jefferson",
					Sex:  "male",
					Age:  4,
				},
			},
		},
		{
			Id:           giraffesId,
			Name:         "giraffes",
			Popularity:   4,
			Location:     "NE",
			Availability: []string{"Wednesday", "Thursday", "Tuesday", "Friday"},
			Residents: []Resident{
				{
					Name: "Gracia",
					Sex:  "female",
					Age:  11,
				},
				{
					Name: "Antone",
					Sex:  "male",
					Age:  9,
				},
				{
					Name: "Vicky",
					Sex:  "female",
					Age:  12,
				},
				{
					Name: "Clay",
					Sex:  "male",
					Age:  4,
				},
				{
					Name: "Arron",
					Sex:  "male",
					Age:  7,
				},
				{
					Name: "Bernard",
					Sex:  "male",
					Age:  6,
				},
			},
		},
	},
	Employees: []Employee{
		{
			Id:             nigelId,
			FirstName:      "Nigel",
			LastName:       "Nelson",
			Managers:       []string{burlId, olaId},
			ResponsibleFor: []string{lionId, tigersId},
		},
		{
			Id:        burlId,
			FirstName: "Burl",
			LastName:  "Bethea",
			Managers:  []string{stephanieId},
			ResponsibleFor: []string{
				lionId,
				tigersId,
				bearsId,
				penguinsId},
		},
		{
			Id:        olaId,
			FirstName: "Ola",
			LastName:  "Orloff",
			Managers:  []string{stephanieId},
			ResponsibleFor: []string{
				ottersId,
				frogsId,
				snakesId,
				elephantsId,
			},
		},
		{
			Id:             wilburnId,
			FirstName:      "Wilburn",
			LastName:       "Wishart",
			Managers:       []string{burlId, olaId},
			ResponsibleFor: []string{snakesId, elephantsId},
		},
		{
			Id:        stephanieId,
			FirstName: "Stephanie",
			LastName:  "Strauss",
			Managers:  []string{},
			ResponsibleFor: []string{
				ottersId,
				giraffesId,
			},
		},
		{
			Id:             sharondaId,
			FirstName:      "Sharonda",
			LastName:       "Spry",
			Managers:       []string{burlId, olaId},
			ResponsibleFor: []string{ottersId, frogsId},
		},
		{
			Id:        ardithId,
			FirstName: "Ardith",
			LastName:  "Azevado",
			Managers:  []string{emeryId},
			ResponsibleFor: []string{
				tigersId,
				bearsId,
			},
		},
		{
			Id:        emeryId,
			FirstName: "Emery",
			LastName:  "Elser",
			Managers:  []string{stephanieId},
			ResponsibleFor: []string{
				lionId,
				bearsId,
				elephantsId,
			},
		},
	},
	Hours: Hours{
		Tuesday:   Day{Open: 8, Close: 6},
		Wednesday: Day{Open: 8, Close: 6},
		Thursday:  Day{Open: 10, Close: 8},
		Friday:    Day{Open: 10, Close: 8},
		Saturday:  Day{Open: 8, Close: 10},
		Sunday:    Day{Open: 8, Close: 8},
		Monday:    Day{Open: 0, Close: 0},
	},
	Prices: Prices{
		Adult:  49.99,
		Senior: 24.99,
		Child:  20.99,
	},
}
