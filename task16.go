package main

func partition(array []int, lowestIdx, highestIdx int) ([]int, int) {
	pivot := array[highestIdx]
	i := lowestIdx
	for j := lowestIdx; j < highestIdx; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[highestIdx] = array[highestIdx], array[i]
	return array, i
}

func quickSort(array []int, lowestIdx, highestIdx int) []int {
	if lowestIdx < highestIdx {
		var p int
		array, p = partition(array, lowestIdx, highestIdx)
		array = quickSort(array, lowestIdx, p-1)
		array = quickSort(array, p+1, highestIdx)
	}
	return array
}

func task16(array []int) []int {
	// sort.Ints тоже используется quicksort )
	return quickSort(array, 0, len(array)-1)
}
