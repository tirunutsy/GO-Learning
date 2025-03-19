package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func main() {
	arr := [] rune {'A','B','C','D'}
	var str string
	for _, c := range arr{
			str += string(c)
	}
	fmt.Println(str)
}
