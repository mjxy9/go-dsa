package main

import "fmt"

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
	DeleteAtBegining(&head)
	DeleteAtBegining(&head)
	Traversal(&head)
}
