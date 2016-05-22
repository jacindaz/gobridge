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

    another_variable := 8
    another_variable_ptr := &another_variable // memory address of 8

    ptr_to_ptr := &another_variable_ptr
    fmt.Println("pointer to a pointer: ", &ptr_to_ptr)
}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func go_by_example_pointer_code() {
    i := 1
    fmt.Println("initial for i:", i)
    zeroval(i)
    fmt.Println("zeroval for i:", i)
    zeroptr(&i)
    fmt.Println("zeroptr for i:", i)
    fmt.Println("pointer for i:", &i)
    fmt.Println(reflect.TypeOf(&i))

    i = 2
    fmt.Println("zeroptr for i:", i)
    fmt.Println("pointer for i:", &i)

    j := 2
    zeroptr(&j)
    fmt.Println("zeroptr for j:", j)
    fmt.Println("pointer for j:", &j)
}
