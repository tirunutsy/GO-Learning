package main

import ( 
"fmt"

)
// LEETCODE NO.9 : Given an integer x, return true if x is a palindrome, and // false otherwise.



func isPalindromeNum(num int) bool {
	
	sNum := string(num)
	length := len(sNum)
	output := ""
	// Following jaggared array
	for i:=0; i<length; i++ {
		for j:=length-1; j>=0; j-- {
		
			if string(sNum[i]) == string(sNum[j]) {
				output += string(sNum[j])
				break;
			}
		}
	}
	fmt.Println(sNum," ",output)
	if sNum == output {
		return true
	}
	
	return false
}

func main() {

	input := 121
	fmt.Println(isPalindromeNum(input))
}
