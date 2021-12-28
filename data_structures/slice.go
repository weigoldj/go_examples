package main

import (
	"fmt"
)

func main() {
	var x []int

	// Remember slices manage arrays, you need to
	// set your variable to the new array.
	// when len == capacity then cap * 2
	fmt.Println(x, len(x), cap(x))

	// append to the slice
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x)) // => 1, 1
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x)) // => 2, 2
	x = append(x, 3)
	fmt.Println(x, len(x), cap(x)) // => 3, 4
	x = append(x, 4)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 5)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 6)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 7)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 8)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 9)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
}
