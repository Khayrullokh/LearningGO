package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Print("-------------------------------\n")
    
	var i int = 0
	for i <= a {
		i = i + 1
		var b int = 0
		for b <= a {
			b = b + 1 
			if i == 0 || i == a-1 || b == 0 || b == a-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
		} 
		fmt.Println()
		
		}
	}
}
