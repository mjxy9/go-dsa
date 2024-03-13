package main

import (
	"fmt"
	"log"
)

type Node struct {
	Data string
	Next *Node
}

func InsertAtBegining(head **Node, data string) {
	
	node := &Node{
		Data: data,
	}

	node.Next = *head
	*head = node
}

func InsertAtEnd(head **Node, data string) {
	
	node := &Node{Data: data}
	temp := *head

	for temp.Next != nil {
		
		temp = temp.Next
	}

	temp.Next = node
}

func InsertAtMiddle(head **Node, data string, position int) {
	
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
		fmt.Printf("[%s]\n", temp.Data)
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

func Search(head **Node, key string) bool {
	
	temp := *head

	for temp != nil {
		
		if temp.Data == key {

			return true
		}
		temp = temp.Next
	}

	return false
}

func main() {
	
	commands := []string{
		"rm -rf main.zip",
		"clear",
		"cd ..",
		"rpm -qa | grep name",
		"cd /usr/lib",
		"touch main.go",
	}

	head := &Node{Data: "code ."}

	for _, command := range commands {
		InsertAtEnd(&head, command)
	}
	
	InsertAtMiddle(&head, "zip -r .", 4)
	Traversal(&head)
}
