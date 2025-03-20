package main

import ( 
"fmt"
)

func startsWith(char string, str string) bool {
	if char == string(str[0]) {
		return true
	}
	return false
}

func main() {
	var c rune = 'A'
	conv := string(c)
	s := "Hello"
	
	fmt.Println(startsWith(conv, s))
}

// OUTPUT : false
