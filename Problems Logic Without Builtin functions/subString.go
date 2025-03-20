package main

import ( 
"fmt"

)

//sub string : A string will have n*(n+1)/2 substrings to achive it we use two loop jaggered array approach where ever


func subString(str string) {
	// We will be giving simply a substring
	for i:=0 ; i<len(str)-1; i++ {
		fmt.Print(string(str[i]))
	}
}

func main() {
	
	str := "Hello"
	
	subString(str)
	//fmt.Println("Type of Output :", reflect.TypeOf(output))
}

// OUTPUT : Hell

