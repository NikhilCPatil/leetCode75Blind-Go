package main

import "fmt"

func containsDuplicate(nums []int) bool {
    
    duplicateMaps := make(map[int]int)

    for _, num := range nums {
        if _,ok := duplicateMaps[num]; ok  {
            return true
        }else {
            duplicateMaps[num] = 1
        }
    }

    return false
}

func main(){
	nums := []int{5,3,2,4,4}
	result := containsDuplicate(nums)
	fmt.Println("Contains duplicates", result);
}