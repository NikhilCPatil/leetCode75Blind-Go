package main

import "fmt"

 type Node struct {
	data int
	next *Node
}

func printLinkedList(head *Node){
	current := head
	for current != nil {
		fmt.Printf("%d ->",current.data)
		current = current.next
	}
}

func reverseLinklist(head *Node)*Node{
var prev, next *Node
current := head

for current != nil {
	next = current.next
	current.next = prev
	prev = current
	current = next
}

return prev
}

func main(){
	head := &Node{ data : 1, next : &Node{ data: 2, next : &Node{data: 3, next: &Node{data: 4, next: nil}} } }
	printLinkedList(head)

	reversedList := reverseLinklist(head)
	printLinkedList(reversedList)
}