package main

func QuickSort(l, h int) {
	if l < h {
		j := Partition(l, h)
		QuickSort(l, j)
		QuickSort(j+1, h)
	}
}

func Partition(l, h int) int {
	pivot := A[l]
	i, j := l, h

	for i < j {
		for A[i] <= pivot {
			i++
		}

		for A[j] > pivot {
			j--
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
		}
	}
	A[l], A[j] = A[j], A[l]
	return j
}
