package main

import "fmt"

func newLoop(x int){
	for i:= 1; i<=10; i++{
	fmt.Println(x*i)
}
}

func loopType2(x int){
	
	for ; x < 1000; {
	x += x
	fmt.Println(x)
}
}

// Using for as while in Go as there is no while in Go
func forAsWhile(x int){
	for x<=10 {
	fmt.Println(x)
	x++
	}
}

func main(){
	a := 0
	for i:=0; i<=10; i++{
		a += i
}
	fmt.Println(a)
	
	b := 5
	newLoop(b)

	c := 1
	loopType2(c)
	
	forAsWhile(1)
}