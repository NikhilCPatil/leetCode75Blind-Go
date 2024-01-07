package main

import "fmt"


func sum(a int,b int) int{
carry := 0
	for b != 0 {
		carry = a & b

		a = a ^ b
	
		b = carry << 1
	}
	

	return a
}

func main(){
a := 5
b := 7
result := sum(a, b)
	fmt.Println(result)
}