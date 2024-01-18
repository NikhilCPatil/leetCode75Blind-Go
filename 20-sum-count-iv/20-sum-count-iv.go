package main

import (
		"fmt"
)


func sumCount(nums []int, target int)int{
d := make([]int, target+1) 
d[0] = 1

for i := 1; i <= target; i++ {
	for _, num := range nums {
		if i >= num{
			d[i] += d[i - num]  
		}
	}
}

// fmt.Println(d)
return d[target]
}

func main(){
	target := 5
	nums := []int{1, 2, 3}
	result := sumCount(nums, target)
	fmt.Println(result)
}