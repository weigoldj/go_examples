//#############################################################################
// Program demonstrates a simple menu with options stored in a map.
// When the main program returns the "defer func" close the listener
//#############################################################################
package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Expresso"

	fmt.Println("MENU")
	fmt.Println("----")
	for key, element := range coffees {
		fmt.Println(fmt.Sprintf("%d - %s", key, element))
	}
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		// rune to string
		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}

	fmt.Println("Program Exiting")
}
