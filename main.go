package main

import (
	"fmt"
	"math"
)

// func main() {
// 	fmt.Println("Enter a number of rows : ")
// 	var rows int
// 	fmt.Scanln(&rows)   ////90C triangle or right triangle

// 	for i := 1; i <= rows; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("* ")
//          fmt.Print("\t*")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	fmt.Println("Enter the height of equilateral triangle : ")
// 	var height int
// 	fmt.Scanln(&height)

// 	for i := 1; i <= height; i++ {
// 		for j := 1; j <= height-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		for j := 1; j <= 2*i-1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	var a, b, c float64

	fmt.Print("Enter coefficient a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter coefficient b: ")
	fmt.Scanln(&b)
	fmt.Print("Enter coefficient c: ")
	fmt.Scanln(&c)

	solution1, solution2, imaginary := solveQuadratic(a, b, c)

	if imaginary {
		fmt.Println("The solutions are imaginary.")
	} else {
		fmt.Printf("Solution 1: %.2f\n", solution1)
		fmt.Printf("Solution 2: %.2f\n", solution2)
	}
}

func solveQuadratic(a, b, c float64) (float64, float64, bool) {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return 0, 0, true
	}
	// not mine and didn't understand what is happening there
	solution1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	solution2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	// especially return and bool type
	return solution1, solution2, false
}

// package main

// import "fmt"

// func main() {
//     var num int
//     fmt.Print("Enter a number: ")
//     fmt.Scanln(&num)

//     for i := 0; i < num; i++ {
//         for j := 0; j < num; j++ {
//             fmt.Print("*")
//         }
//         fmt.Println()
//     }
// }
