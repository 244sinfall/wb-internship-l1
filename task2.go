package main

import (
	"fmt"
	"os"
	"sync"
)

func task2(array []int) {
	output := make([]int, len(array), cap(array))
	wg := sync.WaitGroup{}
	for idx, val := range array {
		wg.Add(1)
		go func(idx int, val int) {
			defer wg.Done()
			output[idx] = val * val
		}(idx, val)
	}
	wg.Wait()
	_, err := fmt.Fprintln(os.Stdout, output)
	if err != nil {
		return
	}
}
