package main

func fibGen(n int) int {
	if n < 2 {
		return 1
	}

	arr := make([]int, n+2)
	arr[1] = 0
	arr[2] = 1
	for i := 3; i < n+1; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}

	return arr[n]
}

func fibGenRec(n int) int {
	if n <= 1 {
		return n
	}
	return fibGenRec(n-1) + fibGenRec(n-2)
}
