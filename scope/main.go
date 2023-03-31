package main

import (
	"fmt"
	"scope/packageone"
)

var packageLevelVariable = "package level value"

func main() {
	var one = "One"

	var packageLevelVariable = "this has the same name as packageLevelVariable"

	fmt.Println(one)
	fmt.Println(packageLevelVariable)

	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()
}

func myFunc() {
	var one = "the number one"

	fmt.Println(one)
	fmt.Println(packageLevelVariable)
}
