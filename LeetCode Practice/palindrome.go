package main

import (
	"fmt"
)


func main() {
	x := "madam"
	fmt.Println(x)
	fmt.Println(len(x))
	 newString := ""
	for _, s := range x {
        newString += s
       // fmt.Printf("The index number of %c is %d\n", s, index)
    }
	fmt.Println(newString)
}
