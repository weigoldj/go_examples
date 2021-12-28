//#############################################################################
// Creates a struct to hold events and prints them out.
//#############################################################################
package main

import (
	"os"
	"text/template"
	"time"
)

type event struct {
	Month     string
	Day       int
	Year      int
	Name      string
	DAYOFWEEK string
}

//make(map[key-type]val-type)
var events = make(map[int]event)
var month = [...]string{"January", "February", "March",
	"April", "May", "June", "July",
	"August", "September", "October",
	"November", "December"}

func init() {
	events[0] = event{getDate(12), 11, 2005, "Happy Birthday Quinn!", getDayOfWeek(1)}
	events[1] = event{getDate(10), 02, 2010, "Happy Birthday Mattaus!", getDayOfWeek(2)}
	events[2] = event{getDate(03), 23, 1966, "Happy Birthday Dani!", getDayOfWeek(3)}

}

func main() {

	tmpl, err := template.New("test").Parse("{{.Day}} {{.Month }} - {{.Name }}, you were born on {{.DAYOFWEEK }}")

	if err != nil {
		panic(err)
	}

	for _, element := range events {
		err = tmpl.Execute(os.Stdout, element)
		if err != nil {
			panic(err)
		}
	}
}

// TODO check for number between 0-11
func getDate(x int) string {
	x = x - 1
	return month[x]
}

func getDayOfWeek(x int) string {
	weekday := time.Now().Weekday().String()
	return weekday
}
