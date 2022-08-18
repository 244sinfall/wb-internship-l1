package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func task4(numberOfWorkers int) {
	channel := make(chan int, numberOfWorkers)
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-stopSignal:
				close(channel)
				return
			default:
				channel <- rand.Int()
			}
		}
	}()
	for w := 1; w <= numberOfWorkers; w++ {
		go worker(w, channel)
	}
}

func worker(id int, jobs <-chan int) {
	for {
		msg, ok := <-jobs
		if ok {
			_, _ = io.WriteString(os.Stdout, "Worker #"+strconv.Itoa(id)+" "+strconv.Itoa(msg)+"\n")
		} else {
			fmt.Printf("Worker %d finished gracefully\n", id)
			return
		}
	}
}
