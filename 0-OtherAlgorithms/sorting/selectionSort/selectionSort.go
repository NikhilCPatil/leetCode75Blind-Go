package main

import "fmt"

//best time complexity = O(n^2)
//average time complexity = O(n^2)
//worst time complexity = O(n^2)

func selectionSort(arr []int) []int {
    smallestKey := 0
    temp := 0
    for i := 0; i < len(arr); i++ {
        smallestKey = i
        for j := i; j < len(arr); j++ {
            if arr[smallestKey] >= arr[j] {
                smallestKey = j
            }
        }
        temp = arr[i]
        arr[i] = arr[smallestKey]
        arr[smallestKey] = temp
    }
    return arr
}

func main() {
    arr := []int{5, 2, 4, 6, 1, 3}
    result := selectionSort(arr)
    fmt.Println(result)
}