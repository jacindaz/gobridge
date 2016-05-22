package main
import "fmt"

func main() {
    my_name := "Jacinda"
    main_ptr := &my_name

    some_ptr_func(main_ptr)
    fmt.Println(*main_ptr)

    tv_show := "Gilmore Girls"
    some_str_func(tv_show)
    fmt.Println(tv_show)
}

func some_ptr_func(my_ptr *string) {
    *my_ptr = "some string"
}

func some_str_func(my_var string) {
    my_var = "some string"
}
