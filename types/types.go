package types

type (
	Hours struct {
		Tuesday   Day `json:"tuesday"`
		Wednesday Day `json:"wednesday"`
		Thursday  Day `json:"thursday"`
		Friday    Day `json:"friday"`
		Saturday  Day `json:"saturday"`
		Sunday    Day `json:"sunday"`
		Monday    Day `json:"monday"`
	}

	Prices struct {
		Adult  float64 `json:"adult"`
		Senior float64 `json:"senior"`
		Child  float64 `json:"child"`
	}

	Resident struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
		Age  int    `json:"age"`
	}

	ZooSpecie struct {
		Id           string     `json:"id"`
		Name         string     `json:"name"`
		Popularity   int        `json:"popularity"`
		Location     string     `json:"location"`
		Availability []string   `json:"availability"`
		Residents    []Resident `json:"residents"`
	}

	Employee struct {
		Id             string   `json:"id"`
		FirstName      string   `json:"firstName"`
		LastName       string   `json:"lastName"`
		Managers       []string `json:"managers"`
		ResponsibleFor []string `json:"responsibleFor"`
	}

	Zoologic struct {
		Species   []ZooSpecie `json:"species"`
		Employees []Employee  `json:"employees"`
		Hours     Hours       `json:"hours"`
		Prices    Prices      `json:"prices"`
	}

	Day struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	}

	AnimalCounterInput struct {
		Specie *string `json:"specie"`
		Sex    *string `json:"sex"`
	}

	AnimalCounterOutput struct {
		Specie   string `json:"specie"`
		Quantity int    `json:"quantity"`
	}
)
