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
