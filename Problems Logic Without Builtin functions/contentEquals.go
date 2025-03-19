package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func main() {
	s := "Tech"
	p := "Tech"
	length := len(s)
	count := 0
	for i, _ := range s{
		if s[i] == p[i] {
			count++
		}
	}
	if length == count {
		fmt.Println("Yes it exists")
	} else {
		fmt.Println("Nope its not exists")
	}
}
