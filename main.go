package main

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type DateDiff struct {
	years        int
	months       int
	weeks        int
	days         int
	hours        int
	minutes      int
	seconds      int
	milliseconds int
	nanoseconds  int
}

type DateTime struct {
	time.Time
}

func (t DateTime) Diff(endDate time.Time, units []string) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	runningDiff := DateDiff{}
	if slices.Contains(units, "years") {
		diff, _ := t.DiffYears(endDate)
		runningDiff.years = diff.years
	}
	if slices.Contains(units, "months") {
		diff, _ := t.DiffMonths(endDate)
		runningDiff.months = diff.months
	}
	if slices.Contains(units, "weeks") {
		diff, _ := t.DiffWeeks(endDate)
		runningDiff.weeks = diff.weeks
	}
	if slices.Contains(units, "days") {
		diff, _ := t.DiffDays(endDate)
		runningDiff.days = diff.days
	}
	if slices.Contains(units, "hours") {
		diff, _ := t.DiffHours(endDate)
		runningDiff.hours = diff.hours
	}
	if slices.Contains(units, "minutes") {
		diff, _ := t.DiffMinutes(endDate)
		runningDiff.minutes = diff.minutes
	}
	if slices.Contains(units, "seconds") {
		diff, _ := t.DiffSeconds(endDate)
		runningDiff.seconds = diff.seconds
	}
	if slices.Contains(units, "milliseconds") {
		diff, _ := t.DiffMilliseconds(endDate)
		runningDiff.milliseconds = diff.milliseconds
	}
	if slices.Contains(units, "nanoseconds") {
		diff := int(endDate.Sub(t.Time))
		runningDiff.nanoseconds = diff
	}

	return runningDiff, nil
}

func (t DateTime) DiffYears(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}
	yearDiff := endDate.Year() - t.Year()
	remainder := endDate.AddDate(-yearDiff, 0, 0).Sub(t.Time)

	if remainder < 0 {
		yearDiff--
		remainder = endDate.AddDate(-yearDiff, 0, 0).Sub(t.Time)
	}

	return DateDiff{years: yearDiff, nanoseconds: int(remainder)}, nil
}

func (t DateTime) DiffMonths(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}
	monthDiff := (endDate.Year()-t.Year())*12 + int(endDate.Month()) - int(t.Month())
	remainder := endDate.AddDate(0, -monthDiff, 0).Sub(t.Time)

	if remainder < 0 {
		monthDiff--
		remainder = endDate.AddDate(0, -monthDiff, 0).Sub(t.Time)
	}

	return DateDiff{months: monthDiff, nanoseconds: int(remainder)}, nil
}

func (t DateTime) DiffWeeks(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	hrInNs := 3600000000000
	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - (hrInNs * int(dateSub.Hours()-0.5))

	weekDiff := int(dateSub.Hours()-0.5) / 24 / 7

	return DateDiff{weeks: weekDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffDays(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	hrInNs := 3600000000000
	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - (hrInNs * int(dateSub.Hours()-0.5))

	dayDiff := int(dateSub.Hours()-0.5) / 24

	return DateDiff{days: dayDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffHours(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	hrInNs := 3600000000000
	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - (hrInNs * int(dateSub.Hours()-0.5))

	hourDiff := int(dateSub.Hours() - 0.5)

	return DateDiff{hours: hourDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffMinutes(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	mnInNs := 60000000000
	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - (mnInNs * int(dateSub.Minutes()-0.5))

	minuteDiff := int(dateSub.Minutes() - 0.5)

	return DateDiff{minutes: minuteDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffSeconds(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	secInNs := 1000000000
	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - (secInNs * int(dateSub.Seconds()-0.5))

	secondDiff := int(dateSub.Seconds() - 0.5)

	return DateDiff{seconds: secondDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffMilliseconds(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	msInNs := 1000000
	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - (msInNs * int(dateSub.Milliseconds()))

	millisecondDiff := int(dateSub.Milliseconds())

	return DateDiff{milliseconds: millisecondDiff, nanoseconds: remainder}, nil
}

func main() {
	foobar := time.Date(2026, 9, 10, 12, 0, 0, 0, time.Local)
	bar := DateTime{time.Now()}

	fmt.Println(bar.Diff(foobar, []string{"years", "months", "weeks", "days", "hours", "minutes"}))
}
