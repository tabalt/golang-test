package main

import (
	"fmt"
	"time"
)

const (
	yearMonthFormat = "200601"
)

type timeInterval struct {
	Begin time.Time
	End   time.Time
}

func splitTimeIntervalByMonth1(ti timeInterval) []timeInterval {
	til := []timeInterval{}
	fmt.Println(ti)
	for {
		beginMonth := ti.Begin.Format(yearMonthFormat)
		endMonth := ti.End.Format(yearMonthFormat)
		if beginMonth == endMonth {
			til = append(til, ti)
			break
		} else {
			mid, _ := time.Parse(yearMonthFormat, beginMonth)
			til = append(til, timeInterval{Begin: ti.Begin, End: mid})
			ti = timeInterval{Begin: mid.Add(time.Second), End: ti.End}

			fmt.Println(mid, ti)
		}
	}
	return til
}

func splitTimeIntervalByMonth(ti timeInterval) []timeInterval {
	til := []timeInterval{}
	if ti.Begin.Before(ti.End) {
		for {
			beginMonth := ti.Begin.Format(yearMonthFormat)
			endMonth := ti.End.Format(yearMonthFormat)
			if beginMonth == endMonth {
				til = append(til, ti)
				break
			} else {
				mid, _ := time.Parse(yearMonthFormat, endMonth)
				til = append(til, timeInterval{Begin: mid, End: ti.End})
				ti = timeInterval{Begin: ti.Begin, End: mid.Add(-1 * time.Second)}
			}
		}
	}
	return til
}

func main() {
	var cases = [][]string{
		{"2019-05-10 01:01:01", "2019-05-15 19:10:20"},
		{"2019-04-10 01:01:01", "2019-05-15 19:10:20"},
		{"2019-03-10 01:01:01", "2019-05-15 19:10:20"},
		{"2019-02-10 01:01:01", "2019-05-15 19:10:20"},
	}

	format := "2006-01-02 15:04:05"
	for _, c := range cases {
		begin, _ := time.Parse(format, c[0])
		end, _ := time.Parse(format, c[1])
		ti := timeInterval{begin, end}

		til := splitTimeIntervalByMonth(ti)
		fmt.Println(ti)
		fmt.Println(til)
		fmt.Println("-------")
	}

}
