package main
import "fmt"

func main() {
    my_function_with_all_the_same_types("Hello", "Jacinda", "blah")
    my_func_with_different_types("Jacinda", 26)
}

func my_function_with_all_the_same_types(str1, str2, str3 string) {
    fmt.Println("Here is string 1: " + str1)
    fmt.Println("Here is string 2: " + str2)
    fmt.Println("Here is string 3: " + str3)
}

func my_func_with_different_types(name string, age int) {
    fmt.Printf("My name is %s and I am %d years old", name, age)
}
