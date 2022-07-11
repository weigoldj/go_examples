package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting date examples")

	p := fmt.Println
	layout := "2006-01-02"
	layout2 := "01-02-2006"
	now := time.Now()
	p(now)

	// get day of week from a date
	ava_date := Date(2000, int(time.December), 11)
	p(ava_date.Weekday())

	mattaus_date := Date(2000, int(time.October), 2)
	p(mattaus_date.Weekday())

	johann_date := Date(2000, int(time.August), 9)
	p(johann_date.Weekday())

	dani_date := Date(2000, int(time.March), 23)
	p("Dani ", dani_date.Format(layout), " ", dani_date.Weekday())
	p("Dani ", dani_date.Format(layout2), " ", dani_date.Weekday())

	p("first day of month ", firstDayOfMonth(now).Weekday())
	p("last day of month ", lastDayOfMonth(now).Weekday())

	test_date := Date(2022, 7, 1)
	p("first day of week ", firstDayOfWeek(test_date).Format(layout2))
}

func lastDayOfMonth(t time.Time) time.Time {
	last_day := Date(t.Year(), int(t.Month()), 1).AddDate(0, 1, 0)
	return last_day.AddDate(0, 0, -1)
}

func firstDayOfMonth(t time.Time) time.Time {
	first_day := Date(t.Year(), int(t.Month()), 1)
	return first_day
}


func firstDayOfWeek(t time.Time) time.Time {
	day_of_week := Date(t.Year(), int(t.Month()), int(t.Day()))
	day_of_week = day_of_week.AddDate(0, 0, - int(day_of_week.Weekday()))
	return day_of_week
}


// creates a date with just year, month and day
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
