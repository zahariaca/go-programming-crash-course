package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")
	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	fmt.Printf("Your name is %s. You are %d years old. Your Favorite number is %2f. Dog owner: %t\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter a valid value.")
		} else {
			return userInput
		}

	}
}

func readInt(s string) int {
	for {
		userInput := readString(s)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number.")
		} else {
			return num
		}

	}
}

func readFloat(s string) float64 {
	for {
		userInput := readString(s)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number.")
		} else {
			return num
		}

	}
}

func readBool(s string) bool {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			fmt.Println("Please y/Y or n/N.")
		}

		if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false
		} else {
			fmt.Println("Please type y/Y or n/N.")
		}
	}
}
