package main

import "fmt"

type SimpleQueue struct {
	Top int
	Rear int
	Size int
	Elements []int
}

func NewSimpleQueue(size int) *SimpleQueue {
	
	return &SimpleQueue{
		
		Top: 0,
		Rear: 0,
		Size: size,
		Elements: make([]int, size),
	}
}


func (q * SimpleQueue) Enqueue(element int) {	
	if q.IsFull() {
		panic("queue is full")
	}

	q.Elements[q.Rear] = element
	q.Rear += 1
}

func (q *SimpleQueue) Dequeue() {
	
	if q.IsEmpty() {
		panic("queue is empty")
	}

	q.Elements[q.Top] = 0
	q.Top += 1

	if q.IsEmpty() {
		q.Top = 0
		q.Rear = 0
	}
}

func (q *SimpleQueue) IsFull() bool {
	
	return q.Rear == q.Size
}

func (q *SimpleQueue) IsEmpty() bool {
	
	return q.Rear == q.Top
}

func (q *SimpleQueue) Print() {
	
	for index := q.Top; index < q.Rear; index++ {
		
		fmt.Printf("[%d]", q.Elements[index])
	}
}


