package main

import (
	"fmt"
	"myapp/packageone"
)

var myVar = "package var"

func main() {
	blockVar := "block Var"

	fmt.Println(myVar, blockVar, packageone.ToString())
	packageone.PrintMe(myVar, blockVar)
	myFunc()
}

func myFunc() {
	one := "The Number One"
	fmt.Println(one)
}
