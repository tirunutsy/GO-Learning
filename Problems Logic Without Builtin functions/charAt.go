package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func main() {
	s := "STITCH"
	index := 3
	// Looping thru all the elemnets with index(i) and value in the index(c) so that we can get desired index output
	for i, c := range s{
	
		if i == index {
		fmt.Printf("The char at %d is %c\n", i, c)
		}
	}
}
