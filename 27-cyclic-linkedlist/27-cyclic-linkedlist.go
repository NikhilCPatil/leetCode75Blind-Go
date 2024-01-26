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

func isCyclic(head *Node)bool{

	if(head == nil || head.next == nil){
		return false
	}

	slow := head
	fast := head.next

	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}

		slow = slow.next
		fast =  fast.next.next
	}


return true
}

func main(){
	head := &Node{ data : 1 }
	node1 := &Node{data : 2}
	node2 := &Node{data : 3}
	node3 := &Node{data : -4}

	head.next = node1
	node1.next = node2
	node2.next = node3
	node3.next = nil

	// printLinkedList(head)
	fmt.Println("Is Cyclic ?",isCyclic(head))
}