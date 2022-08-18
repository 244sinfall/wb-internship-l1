package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task5(seconds int) {
	channel := make(chan int)
	interrupter := make(chan bool)
	go func() {
		defer close(interrupter)
		for {
			select {
			case <-interrupter:
				close(channel)
				return
			default:
				number := rand.Int()
				fmt.Printf("Sending %d...\n", number)
				channel <- number
			}
		}
	}()
	go func() {
		for msg := range channel {
			fmt.Printf("Reading %d...\n", msg)
		}
	}()
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("TIMER PASSED.......................................")
	interrupter <- true
}
