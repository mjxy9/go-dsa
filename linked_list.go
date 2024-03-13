package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func Traversal(head **Node) {
	
	temp := *head

	for temp != nil {
		
		fmt.Printf("[%d] ", temp.Data)
		temp = temp.Next
	}
}

func InsertAtBegining(head **Node, data int) {
	
	node := &Node{Data: data}
	node.Next = *head
	*head = node
}

func InsertAtEnd(head **Node, data int) {
	
	node := &Node{Data: data}
	temp := *head
	
	for temp.Next != nil {
		
		temp = temp.Next
	}

	temp.Next = node
}

func InsertAtMiddle(head **Node, data int, position int) {
	
	node := &Node{Data: data}
	temp := *head

	for index := 2; index < position; index++ {
		
		if temp.Next != nil {
			temp = temp.Next
		}
	}

	node.Next = temp.Next
	temp.Next = node
}

func DeleteAtBegining(head **Node) {
	*head = (*head).Next
}

func DeleteAtEnd(head **Node) {
	
	temp := *head

	for temp.Next.Next != nil {
		temp = temp.Next
	}

	temp.Next = nil
}

func DeleteAtMiddle(head **Node, position int) {
	
	temp := *head

	for index := 2; index < position; index++ {
		
		if temp.Next != nil {
			temp = temp.Next
		}
	}

	temp.Next = temp.Next.Next
}

func Search(head **Node, key int) bool {
	
	temp := *head

	for temp.Next != nil {
		
		if temp.Data == key {
			return true
		}

		temp = temp.Next
	} 

	return false
}

func Sort(head **Node) {
	

	if *head == nil {
		return
	} 

	current := *head
	index := &Node{}
	
	for current != nil {
		
		index = current.Next

		for index != nil {
			
			if current.Data > index.Data {
				temp := current.Data
				current.Data = index.Data
				index.Data = temp
			}

			index = index.Next
		}

		current = current.Next
	}
}

