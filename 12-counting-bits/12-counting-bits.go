package main

import "fmt"

func calculatebits(n int)[]int{
result := make([]int, n+1)
for i := 0; i != n + 1; i++ {
	// fmt.Println(i & 1 )
result[i] = result[i>>1] + ( i & 1 )
}

return result
}

func main(){
num := 5
result := calculatebits(num)
fmt.Println(result)
}