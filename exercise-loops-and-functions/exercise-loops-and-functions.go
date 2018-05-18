package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	oldZ := x
	for oldZ-z > 0.000000001 {
		oldZ = z
		z -= ((z*z - x) / (2 * z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(25))
}
