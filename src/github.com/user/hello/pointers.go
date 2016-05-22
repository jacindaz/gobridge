package main
import "fmt"
import "reflect"

func main() {
    x := 5
    x_ptr := &x

    y := "string"
    y_ptr := &y

    fmt.Println("x value: ", x)
    fmt.Println("x pointer: ", x_ptr)
    fmt.Println("x pointer reflection: ", reflect.TypeOf(x_ptr))

    fmt.Println("\ny value: ", y)
    fmt.Println("y pointer: ", y_ptr)
    fmt.Println("y pointer reflection: ", reflect.TypeOf(y_ptr))
}