package main

import "fmt"
//binarry search
func minimum(nums []int) int{
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		fmt.Println(mid)
		if nums[mid] < nums[right] {
			right = mid
		}else {
			left = mid + 1
		}

	}

	return nums[left]
}

func main(){
	nums :=	[]int{ 4, 5, 6, 7, 0, 1, 2}
	result := minimum(nums)
	fmt.Println("Mimimum number is", result)
}