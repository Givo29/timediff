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

1. Running diff calculation as a boolean
2. Multiple units as a string slice

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
