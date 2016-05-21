package main

import "fmt"

func main() {
    for i := 0; i <= 3; i++ {
        var my_string = fmt.Sprintf("verbose for loop: %d", i)
        fmt.Println(my_string)
    }

    j := 0
    for j <= 3 {
        var my_string = fmt.Sprintf("single condition for loop: %d", j)
        fmt.Println(my_string)
        j++
    }
}
