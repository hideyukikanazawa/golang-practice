package main

import (
	"fmt"
)

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		changed := false
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return arr
}

func main() {
	initial := []int{1, 2, 64, 23, 3, 8, 5, 7}
	fmt.Printf("Initial:\t\t%v\n", initial)
	bubblesort(initial)
	fmt.Printf("Final:\t\t\t%v", initial)
}
