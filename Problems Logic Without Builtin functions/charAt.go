package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func main() {
	s := "STITCH"
	index := 3
	for i, c := range s{
	
		if i == index {
		fmt.Printf("The char at %d is %c\n", i, c)
		}
	}
}