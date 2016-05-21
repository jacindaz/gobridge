package main

import "fmt"

func main() {
    var ten, five int = 10, 5
    var true_true, false_false = true, false
    var random_string = "blah blah"

    fmt.Printf("hello, world\n")
    fmt.Println(ten + five)
    fmt.Println("string interpolation of a string: " + random_string)


    fmt.Println("\nPrintln string interpolation of a boolean: %t", true_true)
    fmt.Println("Println string interpolation of a boolean: %v", false_false)

    fmt.Printf("\nPrintf string interpolation of a boolean: %t", true_true)
    fmt.Printf("\nPrintf string interpolation of a boolean: %v", false_false)

}
