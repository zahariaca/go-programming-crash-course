package main

import "fmt"

func main() {
	var whatToSay string
	whatToSay = "Hello world, again"

	shortHandVariable := "Hello world, again from shorthand"

	sayHelloWorld(whatToSay)
	sayHelloWorld(shortHandVariable)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
