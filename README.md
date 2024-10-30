# TimeDiff

A simple package, relying on no external packages, that extends the standard time package to provide powerful datetime diffing functions. Loosely modelled after the diff function from the [Luxon.js](https://moment.github.io/luxon/#/) package.

## Installation

```
go get github.com/givo29/timediff
```

## Data Structures

TimeDiff only implements two data structures.

1. DateTime - DateTime is the struct that extends the standard time.Time struct to implement the new Diff functions.

```go
type DateTime struct {
	time.Time
}
```

2. DateDiff - DateDiff is the data structure that each of the Diff functions return.

```go
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
```

## Initialise new timediff.DateTime object

```go
import (
    "time"
    "github/givo29/timediff"
)

...

// You dan wrap any standard time.Time initialisor to create a timediff.DateTime
startDate := timediff.DateTime{time.Now()}
```

## Multi Unit Diff Function

The multi-unit diff function allows you to specify two extra parameters over the single unit functions.

1. Multiple units as a string slice
2. Enable a running diff calculation as a boolean

### Multiple Units

The Diff function allows you to specify what units to calculate and return. This is especially useful when used in conjuction with the running parameter below.

Note: The order you specify doesn't matter, the function will always calculate from largest to smallest unit.

```go
startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}    // January 1st 2024
endDate := time.Date(2025, 07, 15, 0, 0, 0, 0, time.UTC)                // July 15th 2025

diff, _ := startDate.Diff(endDate, false, []string{"years", "months", "weeks", "days"}) // Only calculate years, months, weeks and days
fmt.Println(diff)

// Output:
// DateDiff{
//   years: 1
//   months: 18
//   weeks: 80
//   days: 561
// }
```

**Allowed Units:**
| Unit |
|--------------|
| years |
| months |
| weeks |
| days |
| hours |
| minutes |
| seconds |
| milliseconds |
| nanoseconds |

### Running Calculation

If the running parameter is false, each unit will be calculated separately. e.g.

```go
startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}    // January 1st 2024
endDate := time.Date(2025, 07, 15, 0, 0, 0, 0, time.UTC)                // July 15th 2025

diff, _ := startDate.Diff(endDate, false, []string{"years", "months", "weeks"})
fmt.Println(diff)

// Output:
// DateDiff{
//   years: 1
//   months: 18
//   weeks: 80
// }
```

If the running parameter is true, starting with the largest unit, the function will calculate subsequent units relative to the remainder of the previous.
For example, in our example below, 2024/01/01 is 1 year, 6 months **and** 2 weeks from 2025/15/07.

```go
startDate := DateTime{time.Date(2024, 01, 01, 0, 0, 0, 0, time.UTC)}    // January 1st 2024
endDate := time.Date(2025, 07, 15, 0, 0, 0, 0, time.UTC)                // July 15th 2025

diff, _ := startDate.Diff(endDate, false, []string{"years", "months", "weeks"})
fmt.Println(diff)

// Output:
// DateDiff{
//   years: 1
//   months: 6
//   weeks: 2
// }
```

## Single Unit Diff Functions

Each of the single unit functions takes an endDate parameter (time.Time) and outputs a timediff.DateDiff structure with the relavent units. Any remainder will be returned as nanoseconds in the timediff.DateDiff struct.

An error will be returned if the end date is before the start date.

### Function List

| Function         |
| ---------------- |
| DiffMilliseconds |
| DiffSeconds      |
| DiffMinutes      |
| DiffHours        |
| DiffDays         |
| DiffWeeks        |
| DiffMonths       |
| DiffYears        |

### Usage

```go
// One day and one hour difference
startDate := DateTime{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
endDate := time.Date(2024, 1, 2, 1, 0, 0, 0, time.UTC)

diff, _ := startDate.DiffDays(endDate)
fmt.Println(diff)

// Output:
// DateDiff{
//   days: 1
//   nanoseconds: 3600000000000
// }
```
