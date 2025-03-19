package main

import(
"fmt"
"math"
)

func sqrt(x float64) string {
if x < 0 {
return sqrt(-x) + "i" 
}
return fmt.Sprint(math.Sqrt(x)) 
}

func pow(x, n, limit float64) float64{
v := math.Pow(x, n)
if v < limit{
return v
}
return limit
} 

func newPow(x, n, lim float64) float64 {
v := math.Pow(x, n)
	if v < lim {
		
		fmt.Println(v)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	
	//fmt.Println(lim)
	return lim
}

func main() {
 x := 6
//If prcatice 1
if x <= 5{
fmt.Println("Yes The value is within 5")
} else {
fmt.Println("The Value is above 5")
}

//If practice 2
if x%2 == 0 {
fmt.Println("The value", x, "is even")
}else{
fmt.Println("The Value", x, "is odd")
}

//If practice 3
sqr := sqrt(-1)
fmt.Println(sqr)

//If practice 4
if x%3 == 0 {
fmt.Println(x," is Devisible by 3")
} else {
fmt.Println(x," is not Devisible by 3")
}

//If else practice 5
fmt.Println(
pow(3, 2, 10),
pow(3, 3, 20),
)

fmt.Println(
	newPow(3, 2, 10),
)
fmt.Println(
	newPow(3, 3, 20),
)

}