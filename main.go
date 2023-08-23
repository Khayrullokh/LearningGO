package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.0
	var b float64 = -1.0
	var c float64 = 4.0
	// ax^2 + bx + c

	var d float64 = b * b - 4.0 * a * c
	fmt.Printf("d = %f\n", d)
	var sqrt_d float64 = math.Sqrt(d)
	fmt.Printf("sqrt_d = %f\n", sqrt_d)
	
	var x1 float64 = (-b + sqrt_d) / (2 * a)
	fmt.Printf("x1 = %f\n", x1)
	var x2 float64 = (-b - sqrt_d) / (2 * a)
	fmt.Printf("x2 = %f\n", x2)
	
	fmt.Printf("x1 = %f, x2 = %f;\n", x1, x2)
}
