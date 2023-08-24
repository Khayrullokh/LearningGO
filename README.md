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

# Homework:
#### Description: 
    1. Make a 4 comparisons with if ... else ... statements



