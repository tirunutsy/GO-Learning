package main

import ( 
"fmt"
"reflect"
)

func main() {

// 1 createing and slicing elemnts 
	slice := [] int {1,2,3,4}
	fmt.Println(slice)
	fmt.Println("Length : ",len(slice))
	fmt.Println("Capacity : ",cap(slice))
	slice1 := slice[1:3]
	fmt.Println(slice1)
	fmt.Println("Length : ",len(slice1))
	fmt.Println("Capacity : ",cap(slice1))
	
// 2 Appending Elements in slices 
   slice1 = append(slice1, 4)
   fmt.Println(slice1)
  // Adding multiple elements at once 
  slice1 = append(slice1, 5,6,7)
	fmt.Println(slice1)
	fmt.Println("Length : ",len(slice1))
	fmt.Println("Capacity : ",cap(slice1))
// see the type now 
fmt.Println(reflect.TypeOf(slice1))

// 3 Create slice elemnts with mmake 
	slice2 := make([]int, 5, 6)
	fmt.Println(slice2)
	fmt.Println("Length : ",len(slice2))
	fmt.Println("Capacity : ",cap(slice2))
	
// 4 append elments in the slice
	slice2 = append(slice2, 10, 20, 30, 40, 50, 60)
	fmt.Println(slice2)
	fmt.Println("Length : ",len(slice2))
	fmt.Println("Capacity : ",cap(slice2))
	
// 5 Merging 2 differnent slices together in 3rd slice 
   slice3 := append(slice1, slice2...)
   	fmt.Println(slice3)
	fmt.Println("Length : ",len(slice3))
	fmt.Println("Capacity : ",cap(slice3))

// 6 using copy function 
   sliceCopy := make([]int, 4)
   copy(sliceCopy, slice)
   fmt.Println("Slice After Copy")
   fmt.Println(sliceCopy)
   fmt.Println("Length : ",len(sliceCopy))
   fmt.Println("Capacity : ",cap(sliceCopy))

   
}
