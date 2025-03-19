package main

import "fmt"

func checkSame(s1 string, s2 string) bool{
	if s1 == s2 {
	return true
	}
	return false
}

func main() {
	str1 := "Hello"
	str2 := "hello"
	fmt.Println(checkSame(str1,str2))
	
}
