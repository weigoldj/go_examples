/*
  Every GO Programe
  - package name
  - Main function no paramerters
  - imports

  To run 'go run hello.go'

  Go is made up of tokens
  [keyword, identifier, constant, string literal or symbol]

  := "assignment operator"
*/
package main

import (
	"fmt"
)

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}

func main() {
	// Two ways to declare variables
	//	var whatToSay string
	//	whatToSay = "Hello World!!"
	whatToSay := "Hello World!!"

	// If a name is Capitialized then it is exported.
	sayHelloWorld(whatToSay)

	// prints all on the same line.
	fmt.Print("This is a print statement")
	fmt.Print("This is another line")

}
