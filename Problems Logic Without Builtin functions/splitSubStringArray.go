package main

import ( 
"fmt"

)
// subtring in a String Array

func main() {
	
	input := "FLIPS"
	//runeArr := [] rune(input)
	strArr := make([] string, len(input)-1)
	// size := len(input)
	//output := subStringArray(input)
	for i:=1; i<len(input); i++ {
		strArr[i-1] = string(input[i])
	}
	
	fmt.Println(strArr)
	
	
	//for _, v := range output {
		//fmt.Print(v)
	//}
	//fmt.Println("Type of Output :", reflect.TypeOf(output))
}

// OUTPUT : [L I P S]

