package main

//Achieve String charAt() Function logic without using BuiltIn Functions 

import "fmt"

func concateTwo(s1 string, s2 string) string{
	return s1+s2
}

func main() {
	str1 := "Stitch"
	str2 := "Technologies"
	newStr := concateTwo(str1, str2)
	fmt.Println(newStr)
}