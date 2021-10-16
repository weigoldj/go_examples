package main

import (
	"fmt"
)

func init() {
	fmt.Println("I run first")
}

func sumNumbers(param []int) {
	sum := 0
	for _, k := range param {
		sum += k
		fmt.Println("add ", k)
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("I am the main function")

	x := []int{10, 20, 30, 40, 50}
	sumNumbers(x)
}
