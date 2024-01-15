package main

import (
		"fmt"
)


func longestSubsequence(nums []int)int{

if len(nums) == 0 {
	return 0
}

d := make([]int, len(nums)) 
d[0] = 1
maxLength := 1

for i := 1; i < len(nums); i++ {
	d[i] = 1
	for j := 0; j < i; j++{
		if nums[i] > nums[j] && d[i] < d[j] + 1 {
			d[i] = d[j] + 1
		}
	} 

	if d[i] > maxLength{
		maxLength = d[i]
	}
}

return maxLength

}



func main(){
	nums := []int{3,4,-1,0,6,2,3}
	result := longestSubsequence(nums)
	fmt.Println(result)
}