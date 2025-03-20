package main

import ( 
"fmt"
)

func lastIndex(char string, s string) {
	index := 0
	for i:=len(s)-1; i>=0; i-- {
		// fmt.Println(string(s[i]))
		// fmt.Println(i)
		if char == string(s[i]) {
			index = i
			break
		}
	}
	fmt.Println(index)	
}

func main() {
	var c rune = 'e'
	char := string(c)
	s := "Hello"
	lastIndex(char, s)
	//fmt.Println(new, s)
	//fmt.Println(string(chars[0]))
}
