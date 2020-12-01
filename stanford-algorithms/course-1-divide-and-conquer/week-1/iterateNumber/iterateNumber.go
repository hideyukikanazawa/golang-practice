package main

import (
	"fmt"
)

func iterate(n int) int {
	result := 0
	for n != 0 {
		result += n % 10
		n = n/10
	}
	return result
}

func main() {
	n := 9753123
	fmt.Printf("The iterated sum of %v is %v", n, iterate(n))
}