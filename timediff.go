package timediff

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

type DateDiff struct {
	Years        int
	Months       int
	Weeks        int
	Days         int
	Hours        int
	Minutes      int
	Seconds      int
	Milliseconds int
	Nanoseconds  int
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
		runningDiff.Years = diff.Years
		if running {
			endDate = endDate.AddDate(-diff.Years, 0, 0)
		}
	}
	if slices.Contains(units, "months") {
		diff, _ := t.DiffMonths(endDate)
		runningDiff.Months = diff.Months
		if running {
			endDate = endDate.AddDate(0, -diff.Months, 0)
		}
	}
	if slices.Contains(units, "weeks") {
		diff, _ := t.DiffWeeks(endDate)
		runningDiff.Weeks = diff.Weeks
		if running {
			endDate = endDate.AddDate(0, 0, -(diff.Weeks * 7))
		}
	}
	if slices.Contains(units, "days") {
		diff, _ := t.DiffDays(endDate)
		runningDiff.Days = diff.Days
		if running {
			endDate = endDate.AddDate(0, 0, -diff.Days)
		}
	}
	if slices.Contains(units, "hours") {
		diff, _ := t.DiffHours(endDate)
		runningDiff.Hours = diff.Hours
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%dh%dns", diff.Hours, diff.Nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "minutes") {
		diff, _ := t.DiffMinutes(endDate)
		runningDiff.Minutes = diff.Minutes
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%dm%dns", diff.Minutes, diff.Nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "seconds") {
		diff, _ := t.DiffSeconds(endDate)
		runningDiff.Seconds = diff.Seconds
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%ds%dns", diff.Seconds, diff.Nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "milliseconds") {
		diff, _ := t.DiffMilliseconds(endDate)
		runningDiff.Milliseconds = diff.Milliseconds
		if running {
			duration, _ := time.ParseDuration(fmt.Sprintf("-%dms%dns", diff.Milliseconds, diff.Nanoseconds))
			endDate = endDate.Add(duration)
		}
	}
	if slices.Contains(units, "nanoseconds") {
		diff := int(endDate.Sub(t.Time))
		runningDiff.Nanoseconds = diff
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

	return DateDiff{Years: yearDiff, Nanoseconds: int(remainder)}, nil
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

	return DateDiff{Months: monthDiff, Nanoseconds: int(remainder)}, nil
}

func (t DateTime) DiffWeeks(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	weekDiff := int(math.Floor(dateSub.Hours())) / 24 / 7
	return DateDiff{Weeks: weekDiff, Nanoseconds: remainder}, nil
}

func (t DateTime) DiffDays(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	
	dayDiff := int(math.Floor(dateSub.Hours())) / 24
	return DateDiff{Days: dayDiff, Nanoseconds: remainder}, nil
}

func (t DateTime) DiffHours(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	hourDiff := int(math.Floor(dateSub.Hours()))
	return DateDiff{Hours: hourDiff, Nanoseconds: remainder}, nil
}

func (t DateTime) DiffMinutes(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	minuteDiff := int(math.Floor(dateSub.Minutes()))
	return DateDiff{Minutes: minuteDiff, Nanoseconds: remainder}, nil
}

func (t DateTime) DiffSeconds(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	secondDiff := int(math.Floor(dateSub.Seconds()))
	return DateDiff{Seconds: secondDiff, Nanoseconds: remainder}, nil
}

func (t DateTime) DiffMilliseconds(endDate time.Time) (DateDiff, error) {
	if endDate.Before(t.Time) {
		return DateDiff{}, errors.New("End date before start date")
	}

	dateSub := endDate.Sub(t.Time)
	remainder := int(dateSub) - int(dateSub.Nanoseconds())

	millisecondDiff := int(dateSub.Milliseconds())
	return DateDiff{Milliseconds: millisecondDiff, Nanoseconds: remainder}, nil
}
