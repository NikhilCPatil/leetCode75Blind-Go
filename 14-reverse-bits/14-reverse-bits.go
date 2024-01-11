package main

import "fmt"

func reverseBits(num uint32)uint32{
var result uint32 
for i := 0; i < 32; i++ {
	// shift left
	result <<= 1

	result |= num & 1

	//shift right
	num >>=  1

}

return result
}

func main(){
num := uint32(43261596)
result := reverseBits(num)
fmt.Printf("Original: %032b\nReversed: %032b\n", num, result)
}