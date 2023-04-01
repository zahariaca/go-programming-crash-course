package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public (or exported)"
var PackageVar = "packageone.PackageVarValue"

func notExported() {

}

func Exported() {

}

func PrintMe(varOne string, varTwo string, varThree string) {
	fmt.Println(varOne, varTwo, varThree)
}
