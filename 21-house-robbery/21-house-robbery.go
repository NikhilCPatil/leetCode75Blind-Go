package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if(n == 0){
		return 0
	}else if( n == 1 ) {
		return nums[0]
	}
 
	d := make([]int, n)
	 d[0] = nums[0]
	 d[1] = max(nums[0], nums[1])
	 for i:= 2; i < len(nums); i++ {
	   d[i] = max( d[i-1], d[i -2] + nums[i])
	 }
	 return d[n-1]
 }
 
 func max(a int, b int)int{
 if a > b {
	 return a
 }
 return b
 }

 func main(){
	nums := []int{1, 2, 3}
	result := rob(nums)
	fmt.Println(result)
 }