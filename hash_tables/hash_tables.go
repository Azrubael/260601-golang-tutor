package main

import "fmt"

// ArraySize -The size of the hash table
const ArraySize = 7

// HashTable - will hold an array
type HashTable struct {
	array [ArraySize] *bucket
}

// bucket - is a linked list 
type bucket struct {
	head *bucketNode
}

// bucketNode - a structure
type bucketNode struct {
	key string
	next *bucketNode
}

// insert

// search

// delete

// Init - will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)
}