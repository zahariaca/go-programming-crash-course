package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready"

func main() {

	rand.Seed(time.Now().UnixMicro())
	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = rand.Intn(8) + 2

	// another way, declare type and name and assign value
	var secondNumber = rand.Intn(8) + 2

	// one step variable: declare name, assign value and let go figure out the type
	subtraction := rand.Intn(8) + 2

	var answer int
	answer = firstNumber*secondNumber - subtraction

	// display a welcome/instruction
	// take them through the games
	// give them the answer
	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber int, secondNumber int, subtraction int, answer int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)

	fmt.Println("The answer is", answer)
}
