package main

import (
	"fmt"
	"math"
	"sync"
)

func task10(temps []float64) map[int][]float64 {
	output := make(map[int][]float64)
	tempSignal := make(chan float64)
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		for temp := range tempSignal {
			key := int(math.Trunc(temp/10) * 10)
			mu.Lock()
			output[key] = append(output[key], temp)
			mu.Unlock()
		}
		fmt.Println("Выходим из for range проверки канала")
		wg.Done()
	}()
	func() {
		for _, temp := range temps {
			tempSignal <- temp
		}
		defer func() {
			fmt.Println("Закрываем сигнал")
			close(tempSignal)
		}()
	}()
	wg.Wait()
	return output
}
