package main

import (
	"fmt"
	"math/big"
)

func supercomputer400IqCalculator(firstOperand *big.Int, secondOperand *big.Int) {
	var summ big.Int
	var sub big.Int
	var multiply big.Int
	var divis big.Int
	summ.Add(firstOperand, secondOperand)
	sub.Sub(firstOperand, secondOperand)
	multiply.Mul(firstOperand, secondOperand)
	divis.Div(firstOperand, secondOperand)
	fmt.Println("Сумма: " + summ.Text(10))
	fmt.Println("Вычитание: " + sub.Text(10))
	fmt.Println("Деление: " + divis.Text(10))
	fmt.Println("Умножение: " + multiply.Text(10))
}

func task22() {
	// С помощью BigInt
	firstWayFirstOperand, _ := new(big.Int).SetString("9223372036854775807000", 10)
	firstWaySecondOperand, _ := new(big.Int).SetString("20368547758000000000", 10)
	supercomputer400IqCalculator(firstWayFirstOperand, firstWaySecondOperand)
	fmt.Println("Макисмальные числа типов int64, и float64 последовательно. Из uint64 можно вытянуть число в 2 раза больше:")
	int64max := 9223372036854775807
	fmt.Println(int64max)
	float64max := 1.7e308
	fmt.Println(float64max)
	// С помощью Float. Тип определился автоматически
	fmt.Println("Реализация с помощью float")
	secondWayFirstOperand := 10e33
	secondWaySecondOperand := 12.2324232e24
	fmt.Printf("Сумма: %f\n", secondWayFirstOperand+secondWaySecondOperand)
	fmt.Printf("Вычитание: %f\n", secondWayFirstOperand-secondWaySecondOperand)
	fmt.Printf("Деление: %f\n", secondWayFirstOperand/secondWaySecondOperand)
	fmt.Printf("Умножение: %f\n", secondWayFirstOperand*secondWaySecondOperand)
}
