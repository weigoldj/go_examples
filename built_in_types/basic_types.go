package main

import (
	"log"
)

// basic types (numbers, strings, boolean)
var myInt int

// effeciency but discouraged
var myInt16 int16
var myInt32 int32
var myInt64 int64

var myUnit uint // positive values

var myFloat32 float32
var myFloat64 float64

func main() {
	myInt = 10
	log.Println(myInt)

	// immutable strings but act mutable.
	myString := "Ava"
	log.Println(myString, &myString)
	myString = "Quinn" //creates a new string
	log.Println(myString, &myString)

	var myBool = true
	log.Println(myBool)
}
