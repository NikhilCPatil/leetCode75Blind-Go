package main

import ( 
		"fmt"
		"sort"
)



func sum3(nums []int)[][]int{
var result [][]int
sort.Ints(nums)

for i := 0; i < len(nums)-2; i++ {

	if i > 0 && nums[i] == nums[i-1]{
		continue
	}

	target := -nums[i]
	left, right := i + 1, len(nums) - 1

	for left < right {

		sum := nums[left] + nums[right]

		if(sum == target){

			result = append(result, []int{nums[i], nums[left], nums[right]})

			left++
			for left < right && nums[left] == nums[left-1] {
				left++
			}

			right--
			for left < right && nums[right] == nums[right + 1] {
				right--
			}

		}else if sum < target {
			left++
		}else{
			right--
		}

	}

}

return result
}

func main(){
nums := []int{-1, 0, 1, 2, -1, -4}
result := sum3(nums)
fmt.Println(result)
}