package main

import "fmt"

var A = []int{10, 5, 8, 9, 3, 6, 15, 12, 16}
var B = []int{10, 5, 8, 9, 3, 6, 15, 12, 16}

func main() {
	var e1 Employee
	e1 = emp(1)
	e1.Print("ugur")

	fmt.Println("A:")
	fmt.Println(A)
	QuickSort(0, 8)
	fmt.Println(A)

	fmt.Println("B:")
	fmt.Println(B)
	fmt.Println(MergeSort(B))

}
