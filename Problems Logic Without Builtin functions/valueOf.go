package main

import ( 
"fmt"
"reflect"
)

func valueAsString(x string) string{
	var con string
	//fmt.Println(x)
	con = string(x)
	return con
}

func main() {
	var c rune = 'A'
	conv := string(c)
	
	output := valueAsString(conv)
	fmt.Println(output)
	fmt.Println("Type of Output :", reflect.TypeOf(output))
}

// OUTPUT : A
			Type of Output : string
