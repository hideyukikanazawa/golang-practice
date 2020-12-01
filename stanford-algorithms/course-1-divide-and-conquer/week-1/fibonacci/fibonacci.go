package main

import (
	"fmt"
)

func recursiveFibonacci(n int) int {
	if n <= 0 {
		return -1
	} else if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
	}
}


func slightlyBetterFibonacci(n int) int {
	f := make([]int, 2)
	f[0], f[1] = 0, 1
	for i := 2; i <= n-1; i++ {
		f = append(f, f[i-1] + f[i-2])
	}
	return f[n-1]
}

func spaceOptimizedFibonacci(n int) (int, Error()) {
	a, b := 0, 1
	if n <= 0 {
		return -1, err
	} else if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		for i:=2; i<=n-1; i++ {
			a, b = b, a + b
		}
		return b
	}
}



func main() {
	n := 7
	fib := spaceOptimizedFibonacci(n)
	fmt.Printf("The %v-th fibonacci number is %v", n, fib)

}
