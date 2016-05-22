package main
import "fmt"

func main() {
    my_function_with_all_the_same_types("Hello", "Jacinda", "blah")
    my_func_with_different_types("Jacinda", 26)

    some_int, some_bool, some_string := multiple_return_values()
    fmt.Printf("\n\nsome_int: %d, some_bool: %t, some_string: %s", some_int, some_bool, some_string)
}

func my_function_with_all_the_same_types(str1, str2, str3 string) {
    fmt.Println("Here is string 1: " + str1)
    fmt.Println("Here is string 2: " + str2)
    fmt.Println("Here is string 3: " + str3)
}

func my_func_with_different_types(name string, age int) {
    fmt.Printf("My name is %s and I am %d years old", name, age)
}

func multiple_return_values() (int, bool, string) {
    return 100, false, "some string"
}
