package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    var num1 int = 4
    var num2 int = 5
    result := add(num1, num2)

    fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, result)
}