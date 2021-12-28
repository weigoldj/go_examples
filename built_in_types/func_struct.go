package main

import (
	"fmt"
)

// var dog Animal

type Animal struct {
	Name          string
	Sound         string
	NumbersOfLegs int
}

// variadic function takes n variables
func addNumbers(num ...int) int {
	total := 0

	// for index, value in range num
	for _, x := range num {
		total = total + x
	}
	return total
}

// define a function to a type
// func <receiver> name
func (a *Animal) says() {
	fmt.Printf("A %s says %s \n", a.Name, a.Sound)
}

func main() {
	fmt.Println("numbers added = ", addNumbers(1, 2, 3, 4))

	dog := Animal{
		Name:          "Dog",
		Sound:         "Woof",
		NumbersOfLegs: 4,
	}

	dog.says()
}
