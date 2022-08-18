package main

func binarySearch(arr []int, toSearch int) (int, bool) {
	lowest, highest := 0, len(arr)-1
	for lowest <= highest {
		middleIdx := (lowest + highest) / 2
		if arr[middleIdx] == toSearch {
			return middleIdx, true
		}
		if middleIdx == 0 && len(arr) == 1 {
			return 0, false
		}
		if arr[middleIdx] < toSearch {
			lowest = middleIdx + 1
		} else {
			highest = middleIdx - 1
		}
	}
	return 0, false
}

func task17(arr []int, toSearch int) bool {
	sortedArray := quickSort(arr, 0, len(arr)-1)
	_, found := binarySearch(sortedArray, toSearch)
	return found
}
