package main

import (
	"fmt"
	"sort"
)


func mergeSort(arr []int) []int {
	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]
	sort.Ints(arr1)
	sort.Ints(arr2)
	result := []int{}
	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] <= arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}
	result = append(result, arr1...)
	result = append(result, arr2...)
	return result


}

func betterMergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr)/2
	return merge(betterMergeSort(arr[:mid]), betterMergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	size, i, j := len(left) + len(right), 0, 0
	result := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left) - 1 && j <= len(right) -1 {
			result[k] = right[j]
			j++
		} else if j > len(right) -1 && i <= len(left) - 1 {
			result[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}



func main() {
	arr := []int{1, 23, 7, 12, 100, 222, 3}
	fmt.Printf("Initial array is\t\t %v\n", arr)
	fmt.Printf("Final array is\t\t\t %v\n", betterMergeSort(arr))
	fmt.Printf("The half length of array is %v", len(arr)/2)
}