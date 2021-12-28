package main

import (
	"log"
)

// struct considered aggregate type
type car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	year          int
}

func main() {
	// arrays
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	for i, s := range myStrings {
		log.Println(i, s)
	}

	// struct
	myCar := car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Model:         "Mustang",
		Make:          "Ford",
		year:          2001,
	}
	log.Println(myCar)
}
