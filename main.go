package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.0
	var b float64 = 2.0
	var c float64 = 1.0
	// ax^2 + bx + c

	var d float64 = b*b - 4.0*a*c
	fmt.Printf("d = %f\n", d)

	if d > 0 {
		fmt.Println("There are 2 answers")
		var sqrt_d float64 = math.Sqrt(d)
		fmt.Printf("sqrt_d = %f\n", sqrt_d)
		var x1 float64 = (-b + sqrt_d) / (2 * a)
		var x2 float64 = (-b - sqrt_d) / (2 * a)
		fmt.Printf("x1 = %f, x2 = %f;\n", x1, x2)
	} else if d == 0 {
		fmt.Println("There is only one answer")
		x := (-b) / (2 * a)
		fmt.Printf("x = %f;\n", x)
	} else {
		fmt.Println("There is not an answer")
	}
}
