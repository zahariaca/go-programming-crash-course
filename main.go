package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Print("This is some text")
	fmt.Print("This is some  more text")
	sayHelloWorld("")
	sayHelloWorld("Hello World, again!")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
