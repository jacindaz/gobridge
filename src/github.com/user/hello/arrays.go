package main

import "fmt"

func main() {
    my_array := [10]int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(my_array)
    var my_booleans [10]bool

    for i := 0; i < 10; i++ {
        if my_array[i] % 2 == 0 {
            my_booleans[i] = true
        } else {
            my_booleans[i] = false
        }
    }

    fmt.Println(my_booleans)
}
