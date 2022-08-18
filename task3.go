package main

import "sync"

func task3(array []int) int {
	var output int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, val := range array {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			output += val * val
			mu.Unlock()
		}(val)
	}
	wg.Wait()
	return output
}
