package main

import "fmt"

func sort(arr []int)[]int{
	key,j:= 0,0
	for i := 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j=j-1			
		}
		arr[j+1] = key
	}
//   fmt.Println(arr)
	return arr
}

func main(){
	arr := []int{5,2,4,6,1,3}
	result := sort(arr);
	fmt.Println(result)
}