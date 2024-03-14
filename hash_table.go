package main

import (
	"fmt"
	"math"
)

type HashTable struct {
	Array    []Set
	Capacity int
	Size     int
}

type Set struct {
	Key  int
	Data int
}

func NewHashTable(capacity int) *HashTable {

	return &HashTable{
		Array:    make([]Set, capacity),
		Capacity: capacity,
		Size:     0,
	}
}

func (ht *HashTable) Insert(key, data int) {

	index := ht.HashByMultiplication(key)

	if ht.Array[index].Data == 0 {
		ht.Array[index].Key = key
		ht.Array[index].Data = data
		ht.Size += 1
	} else if ht.Array[index].Key == key {
		ht.Array[index].Data = data
	}
}

func (ht *HashTable) Remove(key int) {

	index := ht.HashByMultiplication(key)

	if ht.Array[index].Key == key {
		ht.Array[index].Key = 0
		ht.Array[index].Data = 0
		ht.Size -= 1
	}
}

func (ht *HashTable) Print() {

	for index := 0; index < ht.Capacity; index++ {

		fmt.Printf("[%d] = [%d]\n", ht.Array[index].Key, ht.Array[index].Data)
	}
}

func (ht *HashTable) HashByMultiplication(key int) int {

	fKey := float64(key)
	alpha := (math.Sqrt(5) - 1.0) / 2.0

	return int(math.Floor(float64(ht.Capacity) * math.Mod(fKey*alpha, 1.0)))
}

func main() {

	capacity := 10
	hashTable := NewHashTable(capacity)

	hashTable.Insert(1, 100)
	hashTable.Insert(2, 200)
	hashTable.Insert(3, 300)
	hashTable.Insert(4, 400)
	hashTable.Insert(5, 500)
	hashTable.Insert(6, 600)
	hashTable.Insert(7, 700)
	hashTable.Insert(8, 800)
	hashTable.Insert(9, 900)
	hashTable.Insert(10, 1000)

	hashTable.Remove(1)
	hashTable.Remove(10)

	hashTable.Print()
}
