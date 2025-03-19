package main

import "fmt"

func main() {
// 1 declaring array difffernet ways 
	var arr = [4] int {1,2,3,4}
	arr1 := [3] string {"john", "Paul", "Smith"}
	fmt.Println(arr)
	fmt.Println(arr1)
	
// 2 declaring array difffernet way without any size
    var arr2 = [...]int {1,2,3,4,5}
	arr3 := [5]string {"chennai", "bangalore", "hyd", "mumbai"}
	fmt.Println(arr2)
	fmt.Println(arr3)
	
//3 append values in array in normal way
	arr3[4] = "kochi"
	fmt.Println(arr3)
	
//4 Looping thru elements in array 
	for i:=0; i<len(arr3); i++ {
	fmt.Println(arr3[i])
	}
	
//5 inserting elements in array using loop
	index := 6
	for i:= 0; i<len(arr2); i++ {
	  arr2[i] = index
	  index++
	}
	fmt.Println(arr2)
	
// 6 Create 2d array and Print it
	var arr4 = [][] int {
		{1,1,1},
		{1,1,1},
		{1,1,1},
	}
	fmt.Println(arr4)
	
// 7 Loop thru 2d array and Print it 
	for i:=0; i<len(arr4); i++ {
		fmt.Println(arr4[i])
	}
	
// 8 testing if we can upload elemnts in 2d array by index
  arr4[0][0] = 0
  arr4[1][1] = 0
  arr4[2][2] = 0
  fmt.Println("After changing Values")
	for i:=0; i<len(arr4); i++ {
		fmt.Println(arr4[i])
	}
}
