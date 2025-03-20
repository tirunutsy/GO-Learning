package main

import ( 
"fmt"

)
// Subsequence : loop twice and compare every element based on its ascii value then arrange them based on their asscci value order to get subsequence
// Or We can just print a sub sequqnce travelling from index 1 to len(string)


func subSequence (str string) {
	// We will be giving simply a sub sequence
	for i:=1 ; i<len(str); i++ {
		fmt.Print(string(str[i]))
	}
}

func main() {
	
	input := "BLESS"
	
	subSequence (input)
	//fmt.Println("Type of Output :", reflect.TypeOf(output))
}

// OUTPUT : LESS

