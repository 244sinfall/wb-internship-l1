package main

import (
	"fmt"
	"math/rand"
)

// Убрать глобальную переменную, которая никогда не будет выброшена из памяти
//var justString string

func createHugeString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func someFunc() string {
	// Не генерировать лишнее. Сразу сгенерировать строку нужного размера
	// Остальные 900 символов все равно потеряются после выхода из области действия функции
	//v := createHugeString(1 << 10)
	//justString = v[:100]
	// Вариант 1:
	okString := createHugeString(100)
	return okString
}

func someFunc2() string {
	// Вариант 2:
	okHugeString := createHugeString(1 << 10)
	return okHugeString
}

func task15() {
	justString := someFunc()
	justHugeString := someFunc2()
	// Оставшиеся 900 символов все еще у нас
	justString2 := justHugeString[:100]
	fmt.Println(justString)
	fmt.Println(justString2)
}
