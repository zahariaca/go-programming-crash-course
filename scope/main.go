package main

import (
	"fmt"
	"scope/packageone"
)

var packageLevelVariable = "package level value"
var myVar = "myVarValue"

func main() {
	var one = "One"

	var packageLevelVariable = "this has the same name as packageLevelVariable"

	fmt.Println(one)
	fmt.Println(packageLevelVariable)

	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()

	// variables

	// declare a package level variable for main
	// package named myVar

	// declare a block variable for the main function
	// called blockVar
	var blockVar = "blockVarValue"

	// declare a package level variable in the packageone
	// package named PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function, print out hte values of myVar,
	// blockVar, and PackageVar on oneline using the PrintMe
	// function in packageone
	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}

func myFunc() {
	var one = "the number one"

	fmt.Println(one)
	fmt.Println(packageLevelVariable)
}
