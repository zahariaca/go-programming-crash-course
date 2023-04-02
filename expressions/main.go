package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Righ handed: %t", name, age, rightHanded)
	fmt.Println()

	ageInTenYears := age + 10

	fmt.Printf("In ten years, %s will be %d years old.\n", name, ageInTenYears)

	isATeenager := age >= 13

	fmt.Println(name, "is a teenager:", isATeenager)

	apples := 18
	oranges := 9

	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	fmt.Printf("%d > %d: %t\n", apples, oranges, apples > oranges)
	fmt.Printf("%d < %d: %t\n", apples, oranges, apples < oranges)
	fmt.Printf("%d >= %d: %t\n", apples, oranges, apples >= oranges)
	fmt.Printf("%d <= %d: %t\n", apples, oranges, apples <= oranges)

	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}
	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older.")
		} else {
			fmt.Println(x.Name, "is under 30")

		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 and is over 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 50000 or is under 30")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 50000 or is over 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 50000 and is under 30")
		}
	}
}
