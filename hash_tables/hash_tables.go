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

// Insert - will take a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search - will take a key and return true if the key is stored in the hash table
func (h *HashTable) Search(key string) bool {
		index := hash(key)
		return h.array[index].search(key)
}

// Delete - will take a key and telete the corresponding element from the hash table
// func (h *HashTable) Delete(key string) {
// 		index := hash(key)
// }

// hash - will take a key and return the result of hash function
func hash(key string) int {
	sum := 0
	for _, v := range key{
		sum += int(v)
	}
	return sum % ArraySize
}

// Init - will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// insert - will take a key, kreate a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// search - will take a key and return true if the key is stored in the bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete -

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)
	fmt.Println(hash("RANDY"))
	testHashTable.Insert("RANDY")
	fmt.Println(testHashTable.Search("RANDY"), "RANDY")
	fmt.Println(testHashTable.Search("ANDY"), "ANDY")

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(testBucket.search("RAIN"))
}