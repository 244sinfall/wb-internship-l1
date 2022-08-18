package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type AsyncCounter struct {
	counter int
	mutex   sync.Mutex
}

func (a *AsyncCounter) increment() {
	a.mutex.Lock()
	a.counter++
	a.mutex.Unlock()
}

func task18() *AsyncCounter {
	counter := new(AsyncCounter)
	interrupter := make(chan os.Signal, 1)
	signal.Notify(interrupter, syscall.SIGINT, syscall.SIGTERM)
	func() {
		for {
			select {
			case <-interrupter:
				return
			default:
				go counter.increment()
			}
		}
	}()
	return counter
}
