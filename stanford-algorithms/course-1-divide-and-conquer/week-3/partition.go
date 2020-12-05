package main

import (
	"fmt"
)

func partition(arr []int, l int) []int {
	size := len(arr)
	p := arr[l]
	i := l + 1
	for j := i; j < size; j++ {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i-1], arr[l] = arr[l], arr[i-1]
	return arr
}

func main() {
	pivot := 0
	arr := []int{3, 2, 5, 8, 4, 1, 7}
	fmt.Printf("Initial array: %v", arr)
	partition(arr, pivot)
	fmt.Printf("Final array: %v", arr)
}