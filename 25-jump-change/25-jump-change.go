package main

import "fmt"

func jumpChange(nums []int) bool {
	maxJump := 0
	
	for i := 0; i < len(nums); i++ {
		if i > maxJump{
			return false
		}

		maxJump = max(maxJump, nums[i]+i)

		if(maxJump >= len(nums)-1){
			return true
		}
	}

	   return false
   }

   func max(a int, b int) int{
	if(a > b){
		return a
	}
	return b
   }
   
 func main(){
	nums := []int{2,3,1,1,4}
	result := jumpChange(nums)
	fmt.Println(result)
 }