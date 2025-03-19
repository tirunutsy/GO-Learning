package main

import (
	"fmt"
)

type person struct {
	name string
	age int 
	address string
	salary int
}

func (p person) printPerson() {
	fmt.Println("Name : ", p.name)
	fmt.Println("Age : ", p.age)
	fmt.Println("Address : ", p.address)
	fmt.Println("Salary : ", p.salary)
	fmt.Println()
}

func main() {
	p1 := person{name:"John", age:25, address:"New York", salary:50000}
	
	p1.printPerson()
}
