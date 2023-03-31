package main

type Company struct {
	Name               string
	Location           string
	YearOfFoundation    int
	CurrentYearIncome   int
	PreviousYearIncome  int
	Employees           []string
}
func (c *Company) YearsOfExistence() int {
	return 2023 - c.YearOfFoundation
}

func (c *Company) IncomeDifference() int {
	return c.CurrentYearIncome-c.PreviousYearIncome
}
func (c *Company) HasEmployee(name string) bool {
	for _,v := range c.Employees {
		if v == name {
			return true
		}
	}
	return false
}


