package main

import "fmt"

func jumpChange(nums []int) bool {
	  dp := make([]int, len(nums)+1)
	 fmt.Println(dp)
	   return true
   }
   
 func main(){
	nums := []int{2,3,1,1,4}
	result := jumpChange(nums)
	fmt.Println(result)
 }