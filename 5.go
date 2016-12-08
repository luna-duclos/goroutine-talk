package main

import (
	"log"
	"time"
	"math/rand"
)

func doWork() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

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

	for {
		select {
		case i := <- out1:
			log.Println("1", i)
		case i := <- out2:
			log.Println("2", i)
		}
	}
}
