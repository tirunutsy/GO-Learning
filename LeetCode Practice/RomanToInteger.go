package main

import ( 
"fmt"

)


//func isPalindromeNum(num int) bool {
	
//}

func main() {
	roman := map[string]int {
	"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000,
	}
	//fmt.Println(roman)
	
	str := "XXVI"
	length := len(str)
	var temp string //to store each charAt in loop
	total := 0
	
	for i:=0; i<length; i++ {
		temp = string(str[i])
		for key, value := range roman {
			if temp == key {
				total += value
				break
			}
		}
	}
	fmt.Println("The Given Roman Value in Integer is :", total)
}

// OUTPUT : The Given Roman Value in Integer is : 26

