package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func main() {
	s := "STITCH"
	p := "STITCH"
	length := len(s)
	count := 0
	for i, _ := range s{
		if s[i] == p[i] {
			count++
		}
			
	}
	if length == count {
		fmt.Println("Yes The given String 1:",s ,"And String 2:",p,"Are loxographically correct")
	} else {
		fmt.Println("Nope The given String 1:",s ,"And String 2:",p,"Are not loxographically correct")
	}
}