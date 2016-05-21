package main

import "fmt"

func main() {
    var ten int = 10
    var five int = 5
    var true_true = true
    var random_string = "blah blah"

    fmt.Printf("hello, world\n")
    fmt.Println(ten + five)
    fmt.Println("string interpolation of a string: " + random_string)


    fmt.Println("\nPrintln string interpolation of a boolean: %t", true_true)
    fmt.Println("Println string interpolation of a boolean: %v", true_true)

    fmt.Printf("\nPrintf string interpolation of a boolean: %t", true_true)
    fmt.Printf("\nPrintf string interpolation of a boolean: %v", true_true)

}
