package main

import (
	"log"
	"time"
	"math/rand"
)

func doWork() {
	for i := 0; i < 5; i++ {
		log.Println(i)
		time.Sleep(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)
	}
}

func main() {
	// Just run work in series
	doWork()
	doWork()
}
