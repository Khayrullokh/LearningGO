## Lesson 4: for loop in go and functions intro

```go
func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	multiple := 1
	sum_odd := multiple

	for i := 1; i <= n; i = i + 2 { 
		multiple = i * i * i
		sum_odd = sum_odd + multiple
		fmt.Printf("%d\n", multiple)
	}
	fmt.Printf("Sum - %d\n", sum_odd)
} 
```
```go
package main

import "fmt"

func main() {
	var x float64
	var k float64
	var b float64

	fmt.Print("Enter integers x, k and b : ")
	fmt.Scanln(&x, &k, &b)
	// sum(a, b)
	linear_function(x, k, b)
}

func linear_function(x float64, k float64, b float64) {
	// if b > 0 {
	// 	fmt.Printf("%f * %f + %f = %f", k, x, b, k*x+b)
	// } else if b < 0 {
	// 	fmt.Printf("%f * %f - %f = %f", k, x, -b, k*x+b)
	// } else {
	// 	fmt.Printf("%f * %f = %f", k, x, k*x)
	// }
	
	if b == 0 {
		fmt.Printf("%f * %f = %f", k, x, k*x)
	} else if b < 0 {
		fmt.Printf("%f * %f - %f = %f", k, x, -b, k*x+b)
	} else {
		fmt.Printf("%f * %f + %f = %f", k, x, b, k*x+b)
	}
}

// func sum(a int, b int) {
// 	fmt.Printf("%d + %d = %d ", a, b, a+b)
// }
  

##Revision for loops 
``go
func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	fmt.Print("-------------------------------\n")
	var storage int = b
	var storage_a int = a

	for a <= 10 {
		for b <= 10 {
			fmt.Printf("%d * %d = %d\n", a, b, a * b)
			b = b + 1
		}
		b = storage
		fmt.Println()
		a = a + 1
	}
	a = storage_a
	
	fmt.Printf("a = %d, b = %d ", a, b)
}  