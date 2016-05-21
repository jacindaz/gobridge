package main

import "fmt"

func main() {
    if num := 9; num == 9 {
        fmt.Println("if statement")
        fmt.Printf("this is weird that i can set a variable in the if condition")
    } else {
        fmt.Println("else statement")
    }
}
