package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		x := i + 1
		for j := i + 1; j > -1; j-- {
			if arr[x] < arr[j] {
				arr[x], arr[j] = arr[j], arr[x]
				x = j
			}
		}
	}
	return arr
}
