package main

import (
	"log"
	"time"
	"math/rand"
)

func doWork() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			out <- i
			time.Sleep(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)
		}
	}()

	return out
}

func main() {
	out1 := doWork()
	out2 := doWork()

	for i := range out1 {
		log.Println(i)
	}

	for i := range out2 {
		log.Println(i)
	}
}
