package main

import ( 
"fmt"

)
// LEETCODE NO.1 : Given an array of integers nums and an integer target, return indices of 
// the two numbers such that they add up to target.

func findTargetIndices(arr []int, target int) {
	
	// Following jaggared array
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
		
			if arr[i] + arr[j] == target {
				fmt.Println("The Indices are : [",i,",",j,"]")
				break;
			}
		}
	}
}

func main() {
	arr := []int {2,7,11,15}
	target := 9
	findTargetIndices(arr, target)
}

// OUTPUT : The Indices are : [ 0 , 1 ]

