package main

import (
	"fmt"
	"time"
)

func main() {
	// format
	f := "2006-01-02 15:04:05"
	now := time.Now().Format(f)
	t, _ := time.Parse(f, now)

	fmt.Println("now:\t", now)
	fmt.Println("time:\t", t)
	fmt.Println()

	d := 1 * time.Second
	count := 1

	// timer
	timer := time.NewTimer(d)
	for count <= 3 {
		<-timer.C
		fmt.Println("timer", count)
		count++

		timer.Reset(d)
	}
	fmt.Println()

	// ticker
	ticker := time.NewTicker(d)
	for count <= 6 {
		<-ticker.C

		fmt.Println("ticker", count)
		count++
	}
}
