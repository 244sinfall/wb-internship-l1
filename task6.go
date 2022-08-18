package main

import (
	"fmt"
	"time"
)

// Остановить горутину можно только из нее же.
// Она должна завершиться (прервать свой цикл или выполнится)
// Или получить сигнал остановки (через канал)

func task6() {
	fmt.Println("Создание отдельного канала, по которому потом пройдет сигнал для остановки")
	quit := make(chan bool)
	fmt.Println("Начало горутины...")
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Print(".")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("\nОстанавливаем, давая сигнал в quit")
	quit <- true

}
