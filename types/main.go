package main

import (
	"fmt"
	"sort"

	"github.com/eiannone/keyboard"
)

/*
 basic types (numbers, string, booleans)
*/

var myInt int

// var myInt16 int16 -- exists but shouldn't be used generally
// var myInt32 int32
// var myInt64	int64

var myUint uint // unsigned integer can only hold positive values or 0

var myFloat float32
var myFloat64 float64

var myString string = "John" // string are immutable

var myBool = true

/*
	aggregate types (array, struct)

*/

// structs
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

/*
	refrence types (pointers, slices, maps, functions, channels)
*/

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

var keyPressChan chan rune

/*
	interfaces
*/

type AnimalInterface interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func main() {
	myInt = -10
	myUint = 10

	myFloat = 0.1
	myFloat64 = 100.1
	fmt.Println("-------------------------------------------------")
	// arrays, not generally used, prefer slices
	arraysExample()
	fmt.Println("-------------------------------------------------")
	// structs
	structs()
	fmt.Println("-------------------------------------------------")
	// pointers
	pointers()
	fmt.Println("-------------------------------------------------")
	// slices
	slices()
	fmt.Println("-------------------------------------------------")
	// maps
	maps()
	fmt.Println("-------------------------------------------------")
	// functions
	functions()
	fmt.Println("-------------------------------------------------")
	// channels -- used for go routines
	// goRoutineExample()
	// channels()
	fmt.Println("-------------------------------------------------")
	interfaces()

}

func interfaces() {
	// ask a riddle
	dog := Dog{
		Name: "dog",
		Sound: "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	cat := Cat {
		Name: "Cat",
		Sound: "miau",
		NumberOfLegs: 4,
		HasTail: true,
	}

	Riddle(&cat)
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func Riddle(a AnimalInterface) {
	riddle := fmt.Sprintf(`This animal says "%s" and has "%d" legs. What animal is it?`, d.Sound, d.NumberOfLegs)
	fmt.Println(riddle)
}


func channels() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()

		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

func goRoutineExample() {
	go doSomething("Hello world!")

	fmt.Println("This is another message")
	for {
		// do nothing
	}
}

func doSomething(s string) {
	until := 0
	for {
		fmt.Println("s is", s)
		until = until + 1

		if until >= 5 {
			break
		}
	}

}

func functions() {
	z := addTwoNumbers(2, 4)
	fmt.Println(z)
	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)

	dog := Animal{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "miau",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}

func addTwoNumbers(x, y int) int {
	return x + y
}

// negative return, we can name the return type and do an empty return, not advised
// func addTwoNumbers(x, y int) (sum int) {
// 	sum = x + y
// 	return
// }

// variatic  function, it take a variable number of paramters
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
}

// function with receiver for Animal struct
func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s had %d legs. \n", a.Name, a.NumberOfLegs)
}

func arraysExample() {
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "parrot"
}

func structs() {
	var myFirstCar Car
	myFirstCar.NumberOfTires = 4
	myFirstCar.Luxury = false
	myFirstCar.Make = "Audi"

	mySecondCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC()",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s\n", mySecondCar.Year, mySecondCar.Make, mySecondCar.Model)
}

func maps() {
	intMap := make(map[string]int) // reference type, always passed by reference, and we do not need to use pointers
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// maps are never sorted, and we should never assume that elements are in any given order
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete element "four" from map
	delete(intMap, "four")

	el, ok := intMap["four"]

	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")

	}

}

func slices() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, y := range animals {
		fmt.Println(i, y)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two elements are", animals[0:2]) // [0:2] start at position 0, and give next 2 elements
	fmt.Println("The slice is", len(animals), "elements long")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func pointers() {
	var x int
	x = 10

	fmt.Println("x is", x)

	myFirstPointer := &x
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("After function call, x is now", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
