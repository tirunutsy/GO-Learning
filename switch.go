package main


import "fmt"

func mul(x int){
for i:=1; i<=10; i++{
fmt.Println(x*i)
}
} 

func main(){
a := 3
switch a {
case 1 :
fmt.Println("Its Num 1")
case 2 :
fmt.Println("Its Num 2")
case 3 :
fmt.Println("Its Num 3")
default :
fmt.Println("Running default case")
}

//2 switch
day := 4
switch day{
case 1 :
fmt.Println("Its Sunday")
case 2 :
fmt.Println("Its Monday")
case 3 :
fmt.Println("Its Tuesday")
case 4 :
fmt.Println("Its Wednesday")
case 5 :
fmt.Println("Its Thursday")
case 6 :
fmt.Println("Its Friday")
case 7 :
fmt.Println("Its Saturday")
default:
fmt.Println("Its not any day")
}

x := 2
switch x{
case 1 :
fmt.Println("case 1 running")
case 2 :
fmt.Println("case 2 running")
case 3 :
fmt.Println("case 3 running")
default:
fmt.Println("Enter 1 to 3");
}

//multiple case passing
today := 6
switch today {
case 1,2,3,4,5 :
fmt.Println("Its Weekdays")
case 6,7 :
fmt.Println("Its Weekend")
default:
fmt.Println("Enter correct key");
}

//passing func inside switch
mu := 2
switch mu{
case 1:
mul(1)
case 2:
mul(2)
default:
fmt.Println("Print correct number")
}

}