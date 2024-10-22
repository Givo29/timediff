package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type DateTime struct {
	time.Time
}

func (t DateTime) Diff(endDate time.Time) string {
	runningDate := endDate
	yearsDiff, _, err := t.DiffYears(endDate)
	if err != nil {
		fmt.Println("rip")
		return ""
	}
	runningDate = runningDate.AddDate(-yearsDiff, 0, 0)

	monthsDiff, remainder, err := t.DiffMonths(runningDate)
	if err != nil {
		fmt.Println("rip 2")
		return ""
	}
	runningDate = runningDate.AddDate(0, -monthsDiff, 0)

	return strconv.Itoa(yearsDiff) + " Y, " + strconv.Itoa(monthsDiff) + " M, " + strconv.Itoa(int(remainder))
}

func (t DateTime) DiffYears(endDate time.Time) (int, time.Duration, error) {
	if endDate.Before(t.Time) {
		return 0, 0, errors.New("End date before Start Date")
	}

	yearDiff := endDate.Year() - t.Year()
	remainder := endDate.AddDate(-yearDiff, 0, 0).Sub(t.Time)

	if remainder < 0 {
		yearDiff--
		remainder = endDate.AddDate(-yearDiff, 0, 0).Sub(t.Time)
	}

	return yearDiff, remainder, nil
}

func (t DateTime) DiffMonths(endDate time.Time) (int, time.Duration, error) {
	if endDate.Before(t.Time) {
		return 0, 0, errors.New("End date before Start Date")
	}

	monthDiff := (endDate.Year()-t.Year())*12 + int(endDate.Month()) - int(t.Month())
	remainder := endDate.AddDate(0, -monthDiff, 0).Sub(t.Time)

	if remainder < 0 {
		monthDiff--
		remainder = endDate.AddDate(0, -monthDiff, 0).Sub(t.Time)
	}

	return monthDiff, remainder, nil
}

func main() {
	//	foo := time.Date(2025, 03, 25, 12, 0, 0, 0, time.UTC)
	foobar := time.Date(2026, 9, 10, 12, 0, 0, 0, time.UTC)
	bar := DateTime{time.Now()}

	fmt.Println(bar.Diff(foobar))
}
