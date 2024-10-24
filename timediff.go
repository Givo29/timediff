package timediff

import (
	"errors"
	"fmt"
	"math"
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

func (t DateTime) Diff(endDate time.Time, running bool, units []string) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	runningDiff := DateDiff{}
	if slices.Contains(units, "years") {
		diff, _ := t.DiffYears(endDate)
		runningDiff.years = diff.years
		if running {
			endDate = endDate.AddDate(-diff.years, 0, 0)
		}
	}
	if slices.Contains(units, "months") {
		diff, _ := t.DiffMonths(endDate)
		runningDiff.months = diff.months
		if running {
			endDate = endDate.AddDate(0, -diff.months, 0)
		}
	}
	if slices.Contains(units, "weeks") {
		diff, _ := t.DiffWeeks(endDate)
		runningDiff.weeks = diff.weeks
		if running {
			endDate = endDate.AddDate(0, 0, -(diff.weeks * 7))
		}
	}
	if slices.Contains(units, "days") {
		diff, _ := t.DiffDays(endDate)
		runningDiff.days = diff.days
		if running {
			endDate = endDate.AddDate(0, 0, -diff.days)
		}
	}
	if slices.Contains(units, "hours") {
		diff, _ := t.DiffHours(endDate)
		runningDiff.hours = diff.hours
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%dh%dns", diff.hours, diff.nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "minutes") {
		diff, _ := t.DiffMinutes(endDate)
		runningDiff.minutes = diff.minutes
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%dm%dns", diff.minutes, diff.nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "seconds") {
		diff, _ := t.DiffSeconds(endDate)
		runningDiff.seconds = diff.seconds
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%ds%dns", diff.seconds, diff.nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "milliseconds") {
		diff, _ := t.DiffMilliseconds(endDate)
		runningDiff.milliseconds = diff.milliseconds
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%dms%dns", diff.milliseconds, diff.nanoseconds))
			endDate = endDate.Add(duration)
		}
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

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	weekDiff := int(math.Floor(dateSub.Hours())) / 24 / 7
	return DateDiff{weeks: weekDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffDays(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	
	dayDiff := int(math.Floor(dateSub.Hours())) / 24
	return DateDiff{days: dayDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffHours(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	hourDiff := int(math.Floor(dateSub.Hours()))
	return DateDiff{hours: hourDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffMinutes(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	minuteDiff := int(math.Floor(dateSub.Minutes()))
	return DateDiff{minutes: minuteDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffSeconds(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	secondDiff := int(math.Floor(dateSub.Seconds()))
	return DateDiff{seconds: secondDiff, nanoseconds: remainder}, nil
}

func (t DateTime) DiffMilliseconds(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	millisecondDiff := int(dateSub.Milliseconds())
	return DateDiff{milliseconds: millisecondDiff, nanoseconds: remainder}, nil
}
