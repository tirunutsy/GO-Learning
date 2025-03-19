package main

import "fmt"

func main() {

// Creating a map and assigning values to it normally
	map1 := make(map[int] string)
	map1[1] = "John"
	map1[2] = "Martin"
	map1[3] = "Dhoni"
	map1[4] = "Rohit"
	map1[5] = "Virat"
	fmt.Println(map1)
	
// We can craete and assign maps dynamically as well 
	cars := map[int] string { 1:"BMW", 2:"Lamborghini", 3:"Koinigsegg" }
	fmt.Println(cars)
	
// Access Map elemnts thru loops
	for i:=1; i<len(map1)+1; i++ {
		fmt.Println("Element at", i, "is : ", map1[i])
	}
	
// Accessing and Updating elements at a diferent possitions  
	fmt.Println("Person at 2 :", map1[2])
	
	map1[1] = "Sachin"
	map1[2] = "Bumrah"

    fmt.Println("After Updating the Map :")
	for i:=1; i<len(map1)+1; i++ {
		fmt.Println("Element at", i, "is : ", map1[i])
	}

// Delete an user at a perticular index or key
    
	delete(map1, 1)
	fmt.Println("After Deleting one Element from the Map :")
	for i:=1; i<len(map1)+1; i++ {
		fmt.Println("Element at", i, "is : ", map1[i])
	}


}
