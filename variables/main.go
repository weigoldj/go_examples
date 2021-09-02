package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and dont type your number in, press Enter when ready."

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// declare and assign
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	run(firstNumber, secondNumber, subtraction, answer)
}

func run(firstNumber int, secondNumber int,
	subtraction int, answer int) {

	reader := bufio.NewReader(os.Stdin)
	// display welcome screen
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)

	reader.ReadString('\n')

	// take then through the game
	fmt.Println("Multiple your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiple the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)

	// give the answer
	fmt.Println("The answer is", answer)
}
