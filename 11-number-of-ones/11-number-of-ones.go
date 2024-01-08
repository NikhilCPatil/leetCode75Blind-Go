package main

import "fmt"

func calculateMain(a int)int{
count := 0

for a != 0 {
	count += int(a & 1)
	a >>= 1
}

return count
}

func main(){
num := 5
result := calculateMain(num)
fmt.Println(result)
}