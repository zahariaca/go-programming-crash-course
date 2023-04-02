package staff

import "log"

// accessible outside package
var OverPaidLimit = 75000

// not accessible outside this package
var underPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary > OverPaidLimit {
			overpaid = append(overpaid, x)
		}
	}

	return overpaid
}

func (e *Office) UnderPaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			overpaid = append(overpaid, x)
		}
	}

	return overpaid
}

func (e *Office) notVisible() {
	log.Println("Hello world!")
}
