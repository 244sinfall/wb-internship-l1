package main

import "fmt"

func main() {
	for {
		fmt.Println("Введите номер задания:")
		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Не число! Попробуйте еще раз...")
		}
		switch choice {
		case 1:
			task1()
		case 2:
			task2([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})
		case 3:
			result := task3([]int{2, 4, 6, 8, 10})
			fmt.Println(result)
		case 4:
			task4(40)
		case 5:
			task5(3)
		case 6:
			task6()
		case 7:
			mapToFill := make(map[int]int)
			task7(&mapToFill, 500)
			fmt.Println(mapToFill)
		case 8:
			var initialNumber2 int64 = 59283054
			var initialNumber int64 = 230102512
			task8(&initialNumber2, 8)
			task8(&initialNumber, 8)
		case 9:
			task9([]int{2, 5, 9, 13, 15, 16, 20})
		case 10:
			fmt.Println(task10([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}))
		case 11:
			task11([]string{"dog", "cat", "lion", "tiger"}, []string{"cat", "mouse", "bear", "lion"})
		case 12:
			task12("cat", "cat", "dog", "cat", "tree")
		case 13:
			num1 := 15
			num2 := 30
			fmt.Printf("Число 1 до функции: %d\nЧисло 2 до функции: %d\n", num1, num2)
			task13(&num1, &num2)
			fmt.Printf("Число 1 после функции: %d\nЧисло 2 после функции: %d\n", num1, num2)
		case 14:
			task14(5)
			task14("123")
			task14(true)
			task14(make(chan int))
		case 15:
			task15()
		case 16:
			fmt.Println(task16([]int{3, 7, 8, 2, 4, 5, 6, 1, 20, 3, 459, 230, 22}))
		case 17:
			arr := []int{3, 7, 8, 2, 4, 5, 6, 1, 20, 3, 459, 230, 22}
			toSearch := 459
			fmt.Printf("Ищем %d в %v\n", toSearch, arr)
			found := task17(arr, toSearch)
			if found {
				fmt.Println("Найдено!")
			} else {
				fmt.Println("Не найдено")
			}
		case 18:
			fmt.Printf("Значение счетчика: %d\n", task18().counter)
		case 19:
			str := "главрыба"
			reversedStr := task19(str)
			fmt.Println("Было: " + str + " Стало: " + reversedStr)
		case 20:
			str := "snow dog sun"
			reversedWords := task20(str)
			fmt.Println("Было: " + str + " Стало: " + reversedWords)

		default:
			fmt.Println("Такого задания нет! Попробуйте еще раз...")
		}
	}
}
