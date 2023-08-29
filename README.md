# LearningGO

## Lesson 1: initializing go module
~~~~~~~~~~~~~~~~
Summary:

- go mod init **module_name**
- create new file main.go
- start first lines of code _"package main"_
- func main() { }
- import _"fmt"_
- fmt.Print("description.name\n")
- go build .
- go run . or **module_name**
- complete
~~~~~~~~~~~~~~~~

## Lesson 2:  Variables, fmt.Printf(), and if else
~~~~~~~~~~~~~~~~
Summary:
- how to declare variables in go: 
    1. var <variable_name> <variable_type> = <variable_value>
        e.g. var age int = 18
    2. <variable_name> := <variable_value> => in this case, the compiler will know the type of the variable automatically
        e.g. age := 18
- types:  <int> <uint> <float> <rune> <string> <bool> 
- if <condition> {} else {}:
    if condition is true, then code in if statement is executed. Otherwise, code in else statement is executed.
    in <condition>: && - and, || - or, ! -not
- in fmt.Printf() function: 
    %d - int, uint;
    %f - float;
    %c - rune;
    %s - string;
    %t - bool;

~~~~~~~~~~~~~~~~

## Lesson 3: If else ; fmt.Sprintf(), fmt.Scanln()
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
- if else multiple equations up to 3 comparisons
- fmt.Scanln() <Input by user> :
     e.g. fmt.Scanln(&a, &b)
- fmt.Sprintf() <Makes strings>
     e.g.
     var name string = "Khayrullo"
     var age int = 18
     var full_name string = fmt.Sprintf("I'm %s and I'm %d", name, age)
     
     
- making different equations 
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
## Lesson 4: For loop in go, functions 
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
- for i = 1; i < n; i++ = i + 1 {}  
    e.g for i = 1; i < n; i++
- sum_even or sum_odd 
    e.g sum_even = sum_even + i 
- functions can be added either after or before main function
- func name() -> name cannot be repeated 
- func sum(a int, b int) -> like declaring
    e.g func sum(a int, b int) {  
 	fmt.Printf("%d + %d = %d ", a, b, a+b)
    }

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
## Lesson 5: For loops in go
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
- the basic understanding of for loops 
- how to make quadratic table 
- how to use storages: __WITHOUT CHANGING__ the value it assigns
- how to assign value if there is a , at the end of the function like this :
    e.g : a = a + 1; a = a - 1


~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

## Revision lesson 5: for loops 
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
- how to make rectangular and hollow rectangular using if ... else
  e.g i == 0 || i == a - 1
- how to make square : functions are the same as for retangular
- how to make hollow rectangular with one for loop 
- use i in 2 situations
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


# Homework:
#### Description: 
    1. Fibonacci sequence until n number 
    2. Quadratic equations into functions : a, b, c is given by user and should show the solution numbers as function arguments 
    3. __draw a triangle__ by using if... else ... and for 
    its inside should be filled mostly
    4. Hollow rectangular * 



