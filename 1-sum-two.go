package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // Create a map to store the indices of numbers
    numMap := make(map[int]int)

    // Iterate through the array
    for i, num := range nums {
        // Calculate the complement needed to reach the target
        complement := target - num

        // Check if the complement is in the map
        if index, found := numMap[complement]; found {
			// Return the indices of the two numbers
			fmt.Println(index, i, found, numMap[complement])
            return []int{index, i}
        }

        // If the complement is not in the map, add the current number and its index to the map
        numMap[num] = i
    }

    // If no such pair is found, return an empty slice
    return []int{}
}


func main() {
    nums := []int{5,3,2,4}
    target := 6

    result := twoSum(nums, target)

    if len(result) == 2 {
        fmt.Printf("Indices of the two numbers that add up to %d: %v\n", target, result)
    } else {
        fmt.Println("No such pair found.")
    }
}