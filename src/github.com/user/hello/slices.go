package main

import "fmt"

func main() {
    my_slice := make([][]int, 0)
    fmt.Println(my_slice)

    my_nested_slice := make([]int, 0)
    my_nested_slice = append(my_nested_slice, 34)
    my_nested_slice = append(my_nested_slice, 12)

    my_other_nested_slice := []int{56,78,90}

    my_slice = append(my_slice, my_nested_slice)
    my_slice = append(my_slice, my_other_nested_slice)
    fmt.Println(my_slice)

    my_int_slice := make([]int, 10)
    for i := 0; i < 10; i++ {
        my_int_slice[i] = i + 1
    }

    fmt.Println(my_int_slice)
    fmt.Println(my_int_slice[0:5])
}
