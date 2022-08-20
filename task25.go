package main

import "time"

func letsSleep(d time.Duration) {
	// Функция ждет сигнала, который возвращает time.After. Получение сигнала блокирует выполнение
	<-time.After(d)
}

func task25() {
	letsSleep(5 * time.Second)
}
