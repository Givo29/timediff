package timediff

import (
	"testing"
	"time"
)

// Testing Diff Function
func TestDiffNoRunning(t *testing.T) {
	expectedDiff := DateDiff{years: 1, months: 12, days: 365}

	startDate := DateTime{time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.Diff(endDate, false, []string{"years", "months", "days"})
	if diff != expectedDiff || err != nil {
		t.Fatalf("Diff does not match expected results")
	}
}

func TestDiffNoRunningLeapYear(t *testing.T) {
	expectedDiff := DateDiff{years: 1, months: 12, days: 366}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.Diff(endDate, false, []string{"years", "months", "days"})
	if diff != expectedDiff || err != nil {
		t.Fatalf("Diff does not match expected results")
	}
}

func TestDiffRunning(t *testing.T) {
	expectedDiff := DateDiff{years: 1, months: 6, weeks: 2, hours: 2}
	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2025, 07, 15, 2, 0, 0, 0, time.UTC)

	diff, err := startDate.Diff(endDate, true, []string{"years", "months", "weeks", "hours"})
	if diff != expectedDiff || err != nil {
		t.Fatalf("Diff does not match expected results")
	}
}

// Testing DiffMilliseconds Function
func TestDiffMilliseconds(t *testing.T) {
	const dayInMs = 86400000
	expectedDiff := DateDiff{milliseconds: dayInMs}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffMilliseconds(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffMilliseconds does not match expected results")
	}
}

func TestDiffMillisecondsEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffMilliseconds(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}

// Testing DiffSeconds Function
func TestDiffSeconds(t *testing.T) {
	const dayInSec = 86400
	expectedDiff := DateDiff{seconds: dayInSec}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffSeconds(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffSeconds does not match expected results")
	}
}

func TestDiffSecondsEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffSeconds(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}

// Testing DiffMinutes Function
func TestDiffMinutes(t *testing.T) {
	const dayInMin = 1440
	expectedDiff := DateDiff{minutes: dayInMin}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffMinutes(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffMinutes does not match expected results")
	}
}

func TestDiffMinutesEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffMinutes(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}

// Testing DiffHours Function
func TestDiffHours(t *testing.T) {
	const dayInHr = 24
	expectedDiff := DateDiff{hours: dayInHr}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffHours(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffHours does not match expected results")
	}
}

func TestDiffHoursEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffHours(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}

// Testing DiffDays Function
func TestDiffDays(t *testing.T) {
	const weekInDays = 7
	expectedDiff := DateDiff{days: 7}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 8, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffDays(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffDays does not match expected results")
	}
}

func TestDiffDaysEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffDays(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}

// Testing DiffWeeks Function
func TestDiffWeeks(t *testing.T) {
	const yrInWks = 52
	expectedDiff := DateDiff{weeks: yrInWks}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffWeeks(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffWeeks does not match expected results")
	}
}

func TestDiffWeeksEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffWeeks(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}

// Testing DiffMonths Function
func TestDiffMonths(t *testing.T) {
	const yrInMths = 12 
	expectedDiff := DateDiff{months: yrInMths}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffMonths(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffMonths does not match expected results")
	}
}

func TestDiffMonthsEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffMonths(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}
 
// Testing DiffYears Function
func TestDiffYears(t *testing.T) {
	expectedDiff := DateDiff{years: 10}

	startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2034, 01, 01, 0, 0, 0, 0, time.UTC)

	diff, err := startDate.DiffYears(endDate)
	if diff != expectedDiff || err != nil {
		t.Fatalf("DiffMonths does not match expected results")
	}
}

func TestDiffYearsEarlyEndDate(t *testing.T) {
	startDate := DateTime{time.Date(2024, 01, 02, 0, 0, 0, 0, time.UTC)}
	endDate := time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)

	_, err := startDate.DiffYears(endDate)
	if err == nil {
		t.Fatalf("Expected error but did not receive one")
	}
}
