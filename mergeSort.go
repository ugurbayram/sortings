package main

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	a := MergeSort(arr[:len(arr)/2])
	b := MergeSort(arr[len(arr)/2:])

	return Merge(a, b)

}

func Merge(a, b []int) []int {
	i, j := 0, 0
	var final []int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}
	for j < len(b) {
		final = append(final, b[j])
		j++
	}

	return final

}
