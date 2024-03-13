package main

import (
	"fmt"
	"math"
)

type Set struct {
	key  int
	data int
}

func Insert(array []Set, capacity int, size *int, key int, data int) {

	index := HashByMultiplication(capacity, key)

	if array[index].data == 0 {
		array[index].key = key
		array[index].data = data
		*size += 1
	} else if array[index].key == key {
		array[index].data = data
	}
}

func HashByModule(capacity int, key int) int {
	return capacity % key
}

func HashByMultiplication(capacity int, key int) int {

	alpha := float64(0.458)
	keyFloat := float64(key)

	return int(math.Floor(float64(capacity) * math.Mod((keyFloat*alpha), 1)))
}

func Print(array []Set, capacity int) {

	for index := 0; index < capacity; index++ {
		if array[index].data != 0 {
			fmt.Printf("[%d] = [%d]\n", array[index].key, array[index].data)
		}
	}
}

func main() {

	capacity := 10
	array := make([]Set, capacity)
	size := 0

	Insert(array, capacity, &size, 1, 20)
	Insert(array, capacity, &size, 2, 23)
	Insert(array, capacity, &size, 3, 24)
	Insert(array, capacity, &size, 4, 25)
	Insert(array, capacity, &size, 5, 26)
	Insert(array, capacity, &size, 6, 27)

	Print(array, capacity)

}
