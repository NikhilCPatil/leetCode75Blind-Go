package main

import "fmt"

func missingNum(nums []int)int{
result := len(nums)
for i, num := range nums {
	// fmt.Println(i & 1 )
result ^= i ^ num
}

return result
}

func main(){
nums := []int{3,0,2}
result := missingNum(nums)
fmt.Println(result)
}