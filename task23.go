package main

import "fmt"

// Это решение не изменит массива, находящегося под слайсом
func splice[T any](slice []T, idxWhereToStart, howManyToRemove int) []T {
	return append(slice[:idxWhereToStart], slice[idxWhereToStart+howManyToRemove:]...)
}

// Это решение создаст новый массив и возьмет слайс с него
func remove[T any](slice []T, idxWhereToStart, howManyToRemove int) []T {
	array := make([]T, len(slice)-howManyToRemove)
	k := 0
	for i := 0; i < len(slice); i++ {
		fmt.Println(i, slice[i])
		if i < idxWhereToStart || i >= idxWhereToStart+howManyToRemove {
			array[k] = slice[i]
			k++
		}
	}
	return array
}

func task23() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := array[:]
	slice2 := array2[:]
	slice1RemovedAppendWay := splice(slice1, 4, 2)
	slice2RemovedCopyWay := remove(slice2, 4, 2)
	fmt.Printf("Массив 1: %v\nСлайс 1: %v\nСлайс с удаленными индексами: %v\n\nМассив 2: %v\nСлайс 2: %v\nСлайс с удаленными индексами: %v\n",
		array, slice1, slice1RemovedAppendWay, array2, slice2, slice2RemovedCopyWay)
}
