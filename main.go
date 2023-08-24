package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	
	fmt.Print("Input integer a: ")
	fmt.Scanln(&a)
	fmt.Print("Input integer b: ")
	fmt.Scanln(&b)
	fmt.Print("Input integer c: ")
	fmt.Scanln(&c)

	if a > b && b > c {
		fmt.Printf("%d", a)
	} else if a > c && c > b {
	    fmt.Printf("%d", a)
	} else if b > a && a > c {
		fmt.Printf("%d", b)
	} else if b > c && c > a {
		fmt.Printf("%d", b)
	} else if c > a && a > b {
		fmt.Printf("%d", c)
	} else if c > b && b > a {
		fmt.Printf("%d", c)
	}
}

