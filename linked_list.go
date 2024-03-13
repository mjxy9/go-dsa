package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func InsertAtBegining(head **Node, data int) {
	
	node := &Node{
		Data: data,
	}

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
	
	node := &Node{
		Data: data,
	}
	temp := *head

	for index := 2; index < position; index++ {
		
		if temp.Next != nil {
			temp = temp.Next
		} 
	}

	node.Next = temp.Next
	temp.Next = node
}

func Traversal(head **Node) {
	
	temp := *head

	for temp != nil {
		fmt.Printf("[%d]\n", temp.Data)
		temp = temp.Next
	}
}

func DeleteAtBegining(head **Node) {
	
	*head = ((*head).Next)
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

	for temp != nil {
		
		if temp.Data == key {

			return true
		}
		temp = temp.Next
	}

	return false
}

func Sort(head **Node) {
	
	current := *head
	index := &Node{}
	if current == nil {
		return
	}

	for current != nil {
	
		index = current.Next

		for index != nil {
			
			if current.Data > index.Data {
				aux := current.Data
				current.Data = index.Data
				index.Data = aux
			}

			index = index.Next
		}

		current = current.Next
	}

}

func main() {
	
	data_array := []int{
		11,
		8,
		18,
		19,
		20,
		1,
		200,
		45,
		60,
		6,
	}

	head := &Node{Data: 45}

	for _, data := range data_array {
		InsertAtEnd(&head, data)
	}

	Sort(&head)
	
	Traversal(&head)
}
