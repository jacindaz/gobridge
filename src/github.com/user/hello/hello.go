package main

import "fmt"
import "reflect"
import "errors"
// import "time"

func main() {
    if num := 9; num == 9 {
        fmt.Println("if statement")
        fmt.Printf("this is weird that i can set a variable in the if condition")
    } else {
        fmt.Println("else statement")
    }

    am_i_an_error := true
    the_error := simulating_errors(am_i_an_error)

    var error_to_print = fmt.Errorf("You passed in %t as the argument: ", am_i_an_error)
    fmt.Print(error_to_print)
    fmt.Println(reflect.TypeOf(the_error))
}

func simulating_errors(am_i_an_error bool) error {
    err := errors.New("such corruption, much error")

    if am_i_an_error {
        return err
    } else {
        return nil
    }
}

// func error_formatting(the_error error, am_i_an_error bool) {
//     var error_to_print = fmt.Errorf("You passed in %t as the argument: ", am_i_an_error)
//     fmt.Print(error_to_print)

//     fmt.Println(reflect.TypeOf(the_error))
// }
