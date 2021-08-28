/*
  Arrays is a collection of same elements stored in sequential order.

  Note: a bug in this programs points out that it is hard to add and
  remove objects in array without reindexing the array.

  best to use arrays for fixed data sets.
*/
package main

import (
	"fmt"
)

const maxSize = 50

var (
	studentCount = 0
	students     [maxSize]string
)

func init() {
	addStudent("Ava")
	addStudent("Mattuas")
	addStudent("Otto")
	addStudent("Sandra")
}

func addStudent(name string) {
	if studentCount < maxSize {
		students[studentCount] = name
		studentCount++
	} else {
		fmt.Println("Unable to add any more students")
	}
}

func deleteStudent(name string) {
	for i := 0; i < maxSize; i++ {
		if students[i] == name {
			students[i] = ""
			studentCount--
		}
	}
}

func toString() {
	fmt.Println("Print the Students")
	fmt.Println("class size: ", studentCount)
	for i := 0; i < maxSize; i++ {
		if students[i] != "" {
			fmt.Println("student ", i, ", ", students[i])
		}
	}
}

func main() {
	toString()
	deleteStudent("Sandra")
	toString()
}
