/*
  flag provides command-line options capability.
*/
package main

import (
	"flag"
	"fmt"
)

var (
	envPtr   *string
	portPtr  *int
	msgPtr   *string
	boolPtr  *bool
	greenPtr *bool
	redPtr   *bool

//	color    *Color
)

/*
  While looking thru code i found this eample and liked it
  https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
*/
type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

/*
  Init function are only called once and Go handles this
  implicitly for us.
*/
func init() {
	envPtr = flag.String("env", "development", "current environment")
	portPtr = flag.Int("port", 3000, "port number")
	msgPtr = flag.String("w", "foo", "A String Value")
	boolPtr = flag.Bool("n", false, "a boolean value")
	greenPtr = flag.Bool("ok", false, "display green colorized output")
	redPtr = flag.Bool("error", false, "display green colorized output")
}

func toString() {
	fmt.Println("flag_me values: ")
	fmt.Println("env:", *envPtr)
	fmt.Println("port:", *portPtr)
	fmt.Println("msg:", *msgPtr)
	fmt.Println("bool:", *boolPtr)

	if *greenPtr {
		colorize(ColorGreen, "Nominal")
	} else if *redPtr {
		colorize(ColorRed, "Error: this is in red!")
	} else {
		fmt.Println("no color specified")
	}
}

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	flag.Parse()
	toString()
}
