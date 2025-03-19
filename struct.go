package main

import (
	"fmt"
)

type human struct{
name string
age int 
job string
salary int
}

type vehicle struct{
name string
vehType string
mileage int
year int
}

func printHuman(info human){
fmt.Println("Name : ", info.name)
fmt.Println("Age : ", info.age)
fmt.Println("Job : ", info.job)
fmt.Println("Salary : ", info.salary)
}

func printVehicle(info vehicle){
fmt.Println("Name : ", info.name)
fmt.Println("Type : ", info.vehType)
fmt.Println("Mileage : ", info.mileage)
fmt.Println("Year : ", info.year)
}


func main() {

var human1 human
var human2 human
var human3 human
var vehicle1 vehicle
var vehicle2 vehicle

human1.name = "John"
human1.age = 27
human1.job = "Java Developer"
human1.salary = 50000

human2.name = "Martin"
human2.age = 32
human2.job = "Data Analyst"
human2.salary = 45000

human3.name = "Sam"
human3.age = 25
human3.job = "AI specialist"
human3.salary = 40000

vehicle1.name = "BMW"
vehicle1.vehType = "SUV"
vehicle1.mileage = 30
vehicle1.year = 2022

vehicle2.name = "Lamborgini"
vehicle2.vehType = "Sports"
vehicle2.mileage = 15
vehicle2.year = 2018

// Now call print functions for structs
printHuman(human1)
printHuman(human2)
printHuman(human3)
printVehicle(vehicle1)
printVehicle(vehicle2)

}
