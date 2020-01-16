package main

import (
	"fmt"
	"time"
)

const (
	dbInTimeFormat = "2006-01-02 15:04:05"
)

const (
	dimensionMin  = "10minute"
	dimensionHour = "hour"
	dimensionDay  = "day"
)

func getPointTimeStrList(dimension string, ti timeInterval) []string {

	pointTimeFormat := ""
	stepSecond := time.Duration(0)
	switch dimension {
	case dimensionMin:
		pointTimeFormat = "2006-01-02 15:04:00"
		stepSecond = 10 * 60
	case dimensionHour:
		pointTimeFormat = "2006-01-02 15:00:00"
		stepSecond = 60 * 60
	default:
		pointTimeFormat = "2006-01-02 00:00:00"
		stepSecond = 24 * 60 * 60
	}

	tl := []string{}
	for {
		if ti.Begin.After(ti.End) {
			break
		}
		if dimension == dimensionMin {
			ti.Begin = ti.Begin.Add(time.Duration(-1*ti.Begin.Minute()%10) * time.Minute)
		}

		bt := ti.Begin.Format(pointTimeFormat)
		tl = append(tl, bt)
		ti.Begin = ti.Begin.Add(stepSecond * time.Second)

	}
	if len(tl) > 1 {
		tl = tl[1:]
	}

	return tl
}

func main() {
	var dimensionCases = map[string][][]string{
		"day": {
			{"2019-05-10 01:01:01", "2019-05-15 19:10:20"},
			{"2019-04-10 01:01:01", "2019-05-15 19:10:20"},
			{"2019-03-10 01:01:01", "2019-05-15 19:10:20"},
			{"2019-02-10 01:01:01", "2019-05-15 19:10:20"},
		},
		"hour": {
			{"2019-05-10 01:01:01", "2019-05-10 19:10:20"},
			{"2019-05-10 01:01:01", "2019-05-11 19:10:20"},
			{"2019-05-10 01:01:01", "2019-05-12 19:10:20"},
			{"2019-05-10 01:01:01", "2019-05-13 19:10:20"},
		},
		"10minute": {
			{"2019-05-10 01:01:01", "2019-05-10 01:50:20"},
			{"2019-05-10 01:01:01", "2019-05-10 02:10:20"},
			{"2019-05-10 01:01:01", "2019-05-10 03:10:20"},
			{"2019-05-10 01:01:01", "2019-05-10 04:10:20"},
		},
	}

	format := "2006-01-02 15:04:05"
	for dimension, cases := range dimensionCases {
		for _, c := range cases {
			begin, _ := time.Parse(format, c[0])
			end, _ := time.Parse(format, c[1])
			ti := timeInterval{begin, end}

			ptl := getPointTimeStrList(dimension, ti)

			fmt.Println("-------", dimension, "-------")
			fmt.Println(ti)
			fmt.Println(ptl)

		}
		fmt.Println("\n==============\n")
	}

}
