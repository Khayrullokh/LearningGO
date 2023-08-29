package main

import "fmt"

func main() {
	var width int // 4
	var height int // 5
	fmt.Scanln(&width, &height)
	fmt.Print("-------------------------------\n")

	var i int = 0
	for i < width*height {
		if i%width == 0 && i != 0 {
			fmt.Println()
		}

		if i < width || i > (height - 1) * width || i % width == 0 || i % width == width - 1 {
			fmt.Print("* ")
		} else {
			fmt.Print("  ")
		}
		i = i + 1
	}
}

// var j int = 0
// for j < height {
// 	var i int = 0
// 	for i < width {
// 		fmt.Print("* ")
// 		i = i + 1
// 	}
// 	j = j + 1
// 	fmt.Println()
// }

// var i int = 0
// for i < a {
// 	var b int = 0
// 	for b < a {
// 		if i == 0 || i == a-1 || b == 0 || b == a-1 {
// 			fmt.Print("* ")
// 		} else {
// 			fmt.Print("  ")
// 		}
// 		b = b + 1
// 	}
// 	fmt.Println()
// 	i = i + 1
// }
