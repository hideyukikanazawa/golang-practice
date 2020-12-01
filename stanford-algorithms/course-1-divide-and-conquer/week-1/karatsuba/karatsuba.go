package main

import (
	"fmt"
	"strconv"
	"math"
)

func karatsuba(i, j int) int {
	iLen := len(strconv.Itoa(i))
	jLen := len(strconv.Itoa(j)) 
	
	if iLen == 1 || jLen == 1 {
		return i * j
	}

	x, y := float64(i), float64(j)
	m := math.Max(float64(iLen), float64(jLen))
	m2 := m / 2

	a := int(x / math.Pow(10, m2))
	b := int(x) % int(math.Pow(10, m2))
	c := int(y / math.Pow(10, m2))
	d := int(j) % int(math.Pow(10, m2))

	z0 := karatsuba(b, d)
	z1 := karatsuba(a+b, c+d)
	z2 := karatsuba(a, c)

	return (z2 * int(math.Pow(10, 2*m2))) + ((z1 - z2 - z0) * int(math.Pow(10, m2))) + (z0)
	

}

func main() {
	i := 123123123
	j := 271828182
	final := karatsuba(i, j)
	fmt.Printf("The product of %v and %v is:\t%v", i, j, final)
}
