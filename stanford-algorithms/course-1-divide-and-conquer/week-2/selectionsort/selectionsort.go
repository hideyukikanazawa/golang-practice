package main

import (
	"fmt"
)

func selectionsort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	arr := []int{34, 2, 123, 1, 3, 4, 6, 8}
	fmt.Printf("Initial:\t\t%v\n", arr)
	selectionsort(arr)
	fmt.Printf("Final:\t\t\t%v", arr)
}
