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
	//n = time.Date(n.Year(), 4, 24, n.Hour(), n.Minute(), 00, 000000001, n.Location()) // used for testing
	var l time.Time

	// If is the 25th or past and the salary is paid it moves the date to next month
	switch {
	case n.Day() >= 25:
		l = time.Date(n.Year(), n.Month(), 25, 00, 00, 00, 000000001, n.Location())
		l = l.AddDate(0, 1, 0)
	default:
		l = time.Date(n.Year(), n.Month(), 25, 00, 00, 00, 000000001, n.Location())
	}

	// Runs isTheBackClosed and get the first day before the bank closes
	l = isTheBankClosed(l)

	//Evaluates how to print the output based on close to salary day
	if l.Sub(n).Hours() <= 24 {
		// Prints as x hours/min/sec
		diff := l.Sub(n).String()
		fmt.Printf(diff)
	} else if l.Sub(n).Hours() <= 48 {
		// Prints as x Day
		diff := l.Sub(n).Hours() / 24
		fmt.Printf("%.0f Day!", diff)
	} else {
		// Prints as X Days
		diff := l.Sub(n).Hours() / 24
		fmt.Printf("%.0f Days", diff)
	}
}
