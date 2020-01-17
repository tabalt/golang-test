package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	resultChan := make(chan int, 10)
	n := 10
	step := 1
	for i := 1; i <= n; i += step {
		go partation(ctx, resultChan, i, i+step)
	}

	results := []int{}
	timer := time.NewTimer(2 * time.Second)
loop:
	for {
		select {
		case num := <-resultChan:
			results = append(results, num)
		case <-timer.C:
			log.Printf("main cancel not finished job after 2 seconds\n")
			cancel()
			break loop
		}
	}

	time.Sleep(5 * time.Second)
	log.Printf("all not finished job canceld.\n")
	log.Printf("get %d result: %v\n", len(results), results)
}

func partation(ctx context.Context, results chan int, start, end int) {
	rand.Seed(time.Now().UnixNano())
	secs := rand.Intn(5)
	time.Sleep(time.Duration(secs) * time.Second)

	sum := 0
	for j := start; j < end; j++ {
		select {
		case <-ctx.Done():
			log.Printf("partation %d-%d\tneed %d seconds\tcanceled\n", start, end, secs)
			return
		default:
			sum += j
		}
	}

	log.Printf("partation %d-%d\tneed %d seconds\tcompleted\n", start, end, secs)
	results <- sum
}
