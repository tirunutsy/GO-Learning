package main

import ( 
"fmt"
)

func firstIndex(char string, s string) {
	index := 0
	for i:=0; i<len(s); i++ {
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
	s := "Helleo"
	firstIndex(char, s)
	//fmt.Println(new, s)
	//fmt.Println(string(chars[0]))
}
