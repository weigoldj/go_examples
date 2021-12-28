//#############################################################################
// import a module
// go get -u github.com/eiannone/keyboard
package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

// user defined type
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
	user.OwnsADog = readBool("Do you own a dog?")

	// String interpolation effecient to least effecient
	fmt.Printf(fmt.Sprintf("Your name is %s. You are %d years old. Your favorite number is %.2f and you own a dog %t \n",
		user.UserName, user.Age, user.FavoriteNumber, user.OwnsADog))
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	// fmt.Println("Your name is "+userName+". You are ", age, "years old")
}

func prompt() {
	fmt.Println("-> ")
}

func readString(str string) string {
	for {
		fmt.Println(str)
		prompt()

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput != "" {
			return userInput
		} else {
			str = "What is your name??"
		}
	}
}

func readInt(str string) int {
	for {
		fmt.Println(str)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // linux

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number.")
		} else {
			return num
		}
	}
}

func readFloat(str string) float64 {
	for {
		fmt.Println(str)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // linux
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number.")
		} else {
			return num
		}
	}
}

func readBool(str string) bool {

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(str)
		prompt()
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please enter y or new")
		} else if char == 'y' || char == 'Y' {
			return true
		} else {
			return false
		}
	}
}
