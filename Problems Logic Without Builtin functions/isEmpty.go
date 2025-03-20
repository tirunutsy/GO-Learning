package main

import ( 
"fmt"
)

func isEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func main() {
	s := "a"
	
	fmt.Println(isEmpty(s))
}

// OUTPUT : false