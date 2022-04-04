package main

import (
	"fmt"
	"time"
)

type Q struct {
	Weeks int
	Month time.Month
	Day int
}

var qs = map[int]Q{
	1: {12, time.January, 10},
	2: {13, time.April, 1},
	3: {13, time.July, 1},
	4: {13, time.October, 1},
}

func main() {
	year := 2022
	qN := 2

	fmt.Println(fmt.Sprintf("# %d Q%d", year, qN))
	fmt.Println()

	qStart, weeks := getQStartAndWeeks(year,qN)
	printQ(qStart, weeks)
}

func getQStartAndWeeks(year, qN int) (time.Time, int) {
	q := qs[qN]
	var qStart time.Time
	for i := 0; i < 7; i++ {
		qStart = time.Date(year, q.Month, q.Day+i, 0, 0, 0, 0, time.Local)
		if qStart.Weekday() == time.Monday {
			break
		}
	}
	return qStart, q.Weeks
}

func printQ(qStart time.Time, weeks int) {
	if qStart.Weekday() != time.Monday {
		panic(fmt.Errorf("qStart must be Monday, got: %v", qStart.Weekday()))
	}

	const weekDuration = 7 * 24 * time.Hour

	for i := 0; i < weeks; i++ {
		printWeek(i+1, qStart.Add(time.Duration(i)*weekDuration))
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

func printWeek(weekNum int, weekStart time.Time) {
	fmt.Println(fmt.Sprintf("## %02d неделя (_%s - %s_)", weekNum, weekStart.Format("2006-01-02"), weekStart.Add(4*24*time.Hour).Format("2006-01-02")))
	fmt.Println("### Цели")
	fmt.Println("- ")
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Println(fmt.Sprintf("#### %s:", weekStart.Add(time.Duration(i)*24*time.Hour).Format("2006-01-02")))
		fmt.Println("- ")
		fmt.Println()
	}
	fmt.Println("### summary")
	fmt.Println("- ")
}
