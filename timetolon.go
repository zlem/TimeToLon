package main

import (
	"fmt"
	"time"

	"github.com/zlem/cal"
)

func isTheBankClosed(date time.Time) time.Time {
	// Creates a "Calendar" in cal lib and adds the swedish holidays
	// Checks if date is a holiday and retracts one day until its not a holiday anymore
	c := cal.NewCalendar()
	cal.AddSwedenHolidays(c)

	for c.IsWorkday(date) == false {
		date = date.AddDate(0, 0, -1)
	}
	return date
}

func main() {
	n := time.Now()
	//n = time.Date(n.Year(), 4, 24, 13, n.Minute(), 00, 000000001, n.Location()) // used for testing
	l := isTheBankClosed(time.Date(n.Year(), n.Month(), 25, 00, 00, 00, 000000001, n.Location()))

	// If is the payday or past and the salary is paid it moves the date to next month
	switch {
	case n.Day() >= l.Day():
		l = l.AddDate(0, 1, 0)
	}

	//Evaluates how to print the output based on close to salary day
	if l.Sub(n).Hours() <= 24 {
		// Prints as x hours/min/sec
		diff := l.Sub(n).String()
		fmt.Printf(diff)
	} else {
		// Prints as X Days
		diff := l.Sub(n).Hours() / 24
		fmt.Print(int(diff)+1, " Days")
	}
}
