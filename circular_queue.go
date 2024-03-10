package main

import "fmt"

type Queue struct {
	Top int
	Rear int
	Size int
	Elements []int
}

func NewQueue(size int) *Queue {
	
	return &Queue{
		Top: 0,
		Rear: 0,
		Size: size,
		Elements: make([]int, size),
	}
}

func (q *Queue) Enqueue(element int) {
	
	if q.IsFull() {
		panic("queue is full")
	}

	q.Elements[q.Rear % q.Size] = element
	q.Rear += 1
}

func (q *Queue) Dequeue() {
	
	if q.IsEmpty() {
		panic("queue is empty")
	}

	q.Elements[q.Top % q.Size] = 0
	q.Top += 1
}

func (q *Queue) IsFull() bool {
	
	return (q.Rear - q.Top) == q.Size
}

func (q *Queue) IsEmpty() bool {
	return q.Rear == q.Top
}

func (q *Queue) Print() {
	
	for index := q.Top; index < q.Rear; index++ {
		fmt.Printf("[%d]", q.Elements[index % q.Size])
	}
}

