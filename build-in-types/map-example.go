package main

import (
	"fmt"
)

// maps are keys and values
// to create a map, "make ( map [with key type] to hold type)"
// maps is a reference type and is passed by reference not pointers.
//
func main() {
	// define a map
	myMap := make(map[string]string)

	// add key value pairs
	myMap["IndianRed"] = "CD5C5C"
	myMap["FireBrick"] = "B22222"
	myMap["DeepPink"] = "FF1493"
	myMap["OrangeRed"] = "FF4500"
	myMap["Yellow"] = "FFFF00"

	// print the map
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// delete entry
	delete(myMap, "OrangeRed")

	// does entry exist, ok = bool true if removed
	el, ok := myMap["OrangeRed"]
	if ok {
		fmt.Println(el, "was found")
	} else {
		fmt.Println(el, "was not found")
	}

}
