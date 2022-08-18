package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func task9(from []int) {
	input := make(chan int)
	output := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		defer func() {
			fmt.Println("Выход из горутины, обрабатывающей вход")
			close(output)
		}()
		for inputInt := range input {
			output <- inputInt * 2
		}
	}()
	go func() {
		defer fmt.Println("Выход из горутины, обрабатывающей выход")
		for outputInt := range output {
			outStr := fmt.Sprintf("%d\n", outputInt)
			_, _ = io.WriteString(os.Stdout, outStr)
		}
	}()
	for _, val := range from {
		wg.Add(1)
		go func(val int) {
			input <- val
			wg.Done()
		}(val)
	}
	wg.Wait()
	close(input)
}
