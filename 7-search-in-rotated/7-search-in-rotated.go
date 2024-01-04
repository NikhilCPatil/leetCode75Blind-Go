package main

import "fmt"
//binarry search
func search(nums []int, target int) int{
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(mid)

		if(nums[mid] == target){
			return mid
		}

		if nums[left] <= nums[mid] {

		if(nums[left] <= target && target  < nums[mid] ){
			right = mid - 1
		}else{
			left = mid + 1
		}
		}else {
			if(target <= nums[right]  && nums[mid] < target){
				left = mid + 1
			}else{
				right = mid - 1
			}
		}

	}

	return -1
}

func main(){
	nums :=	[]int{ 4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := search(nums, target)
	fmt.Println("Mimimum number is", result)
}