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

func mergeLinkedlist(list1 *Node, list2 *Node)*Node{
dummy := &Node{next: list1}
current := dummy

for list1 != nil && list2 != nil {
	if list1.data <= list2.data{
		list1 = list1.next
	}else{
		next2 := list2.next
		list2.next = current.next
		current.next = list2
		list2 = next2
	}
	current = current.next
}

if list2 != nil{
	current.next = list2
}

return dummy.next
}

func main(){
	list1 := &Node{ data : 1, next : &Node{ data: 2, next : &Node{data: 3, next: &Node{data: 5, next: nil}} } }
	list2 := &Node{ data : 1, next : &Node{ data: 2, next : &Node{data: 3, next: &Node{data: 4, next: nil}} } }

	merged := mergeLinkedlist(list1, list2)

	printLinkedList(merged)
}