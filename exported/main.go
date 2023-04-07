package main

import (
	"exported/staff"
	"log"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())

	staff.OverPaidLimit = 10000
	log.Println("Overpaid Staff", myStaff.Overpaid())

	// staff.UnderPaidLimit = 10000 <- does not work because variable is not exported outside package (starts with lowercase letter)
	log.Println("Underpaid Staff", myStaff.UnderPaid())

	// myStaff.notVisible - not exported, same as above variable
}
