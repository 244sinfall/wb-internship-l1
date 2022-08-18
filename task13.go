package main

func task13(number1, number2 *int) {
	*number2 = *number1 + *number2
	*number1 = *number2 - *number1
	*number2 = *number2 - *number1
}
