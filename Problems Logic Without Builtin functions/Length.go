package main

import ( 
"fmt"
)

func findLength(str string) int {
	length := 0
	if len(str) == 0 {
		return length 
	} else {
		for i, _ := range str{
			length += i
		}
	}
	
	return length+1
}

func main() {
	s := ""
	
	fmt.Println(findLength(s))
}

// OUTPUT : 0