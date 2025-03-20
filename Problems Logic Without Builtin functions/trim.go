package main

import ( 
"fmt"

)

func trimString(x string) string{
	var trimmed string = ""
	for _, s := range x {
	//fmt.Println(string(s))
		if string(s) == " " {
 			continue
		} else {
			trimmed += string(s)
		}
	}
	return trimmed
}

func main() {
	
	str := "Hi Hello"
	output := trimString(str)
	fmt.Println(output)
	//fmt.Println("Type of Output :", reflect.TypeOf(output))
}

// OUTPUT : HiHello

