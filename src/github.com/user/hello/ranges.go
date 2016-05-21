package main
import "fmt"

func main() {
    my_array := [10]int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(my_array)

    for i, num := range my_array {
        if my_array[i] % 2 == 0 {
            fmt.Printf("\n%d is ", num)
            fmt.Printf("even")
        } else {
            fmt.Printf("\n%d is ", num)
            fmt.Printf("odd")
        }
    }
}
