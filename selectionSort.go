package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if arr[min] != arr[i] {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
