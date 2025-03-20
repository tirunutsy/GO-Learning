package main

import ( 
"fmt"
)

func convertToCharArray(s string) []rune {
	newRune := [] rune (s)
	return newRune
}

func main() {
	//var c rune = 'A'
	s := "Hello"
	output := convertToCharArray(s)
	
	for _, c := range output {
		fmt.Print(string(c)," ")
	}
}

// OUTPUT : false
