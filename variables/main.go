package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready"

func main() {
	
	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 2
	
	// another way, declare type and name and assign value
	var secondNumber = 5
	
	// one step variable: declare name, assign value and let go figure out the type
	subtraction := 7
	
	var answer int
	
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instruction

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the games

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)

	// give them the answer

}
