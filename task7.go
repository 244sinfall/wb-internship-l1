package main

import (
	"math/rand"
	"sync"
)

func task7(to *map[int]int, elementsCount int) {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	for i := 0; i < elementsCount; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			(*to)[i] = rand.Int()
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
