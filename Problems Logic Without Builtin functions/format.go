package main

import ( 
"fmt"
"reflect"
)

func convertString(x int) string{
	var con string
	fmt.Println(x)
	con = string(x)
	return con
}

//For exact String value we have to use fmt.SprintF or "strconv"

func main() {
	x := 55
	s := convertString(x)
	fmt.Println(s)
	fmt.Println(s,"Type is :",reflect.TypeOf(s))
	
}
