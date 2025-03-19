package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func main() {
	
	str := "Hello"
	last := "o"
	var output string
	for i, c := range str{
		if(len(str)-1 == i){
			output += string(c)
		}
	}
	if(last == output){
		fmt.Println("Yes It ends with same")
	} else {
		fmt.Println("No its deferent")
	}
}
