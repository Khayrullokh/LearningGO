package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Print("-------------------------------\n")
    
	var i int = 1
	for i <= a {
		var b int = 1
		for b <= a {
			fmt.Printf("* ")
			b = b + 1
		} 
		fmt.Println()
		i = i + 1
	}
}