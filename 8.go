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
	out3 := doWork()

	for out1 != nil && out2 != nil && out3 != nil {
		select {
		case i, ok := <- out1:
			if !ok {
				out1 = nil
				continue
			}

			log.Println("1", i)
		case i, ok := <- out2:
			if !ok {
				out2 = nil
				continue
			}

			log.Println("2", i)
		case i, ok := <- out3:
			if !ok {
				out3 = nil
				continue
			}

			log.Println("3", i)
		}
	}
}
