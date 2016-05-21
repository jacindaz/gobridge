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

    string_short_var := "short variable"
    fmt.Printf("\n\n%v", string_short_var)

    boolean_short_var := true
    fmt.Printf("\nboolean short var: %t", boolean_short_var)

    var boolean_no_value bool
    fmt.Println("\n")
    fmt.Println(boolean_no_value)

    var int_no_value int
    fmt.Println(int_no_value)

    const my_constant = 100
    fmt.Println("My constant: %s", my_constant)

    my_constant = "trying to change my constant"
    fmt.Println("My constant: %s", my_constant)
}
