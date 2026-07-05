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

// Init - will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
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

// Delete - will take a key and delete the corresponding element from the hash table
func (h *HashTable) Delete(key string) {
		index := hash(key)
		if !h.array[index].search(key) {
			return
		}
		if h.array[index].head.key == key {
			h.array[index].head = h.array[index].head.next
		}
		h.array[index].delete(key)
}

// insert - will take a key, kreate a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "is already in the bucket")
	}
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

// delete - will take a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
		prevNode := b.head
		for prevNode.next != nil {
			if prevNode.next.key == k {
				prevNode.next = prevNode.next.next
				return
			}
			prevNode = prevNode.next
		}
}

// hash - will take a key and return the result of hash function
func hash(key string) int {
	sum := 0
	for _, v := range key{
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)
	fmt.Println(hash("RANDY"))
	testHashTable.Insert("RANDY")
	testHashTable.Insert("RANDY")
	fmt.Println(testHashTable.Search("RANDY"), "RANDY")
	fmt.Println(testHashTable.Search("ANDY"), "ANDY")

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	fmt.Println("RANDY is in the bucket", testBucket.search("RANDY"))
	fmt.Println("RAIN is in the bucket", testBucket.search("RAIN"))
	testBucket.insert("RAIN")
	fmt.Println("RAIN is in the bucket", testBucket.search("RAIN"))
	testBucket.delete("RAIN")
	fmt.Println("RAIN is in the bucket", testBucket.search("RAIN"))

	fmt.Print("\n\n\n")
	ht := Init()
	ll := []string {
		"Vera",
		"Cat",
		"line",
		"ERIC",
		"rainy",
		"rainier",
		"rainiest",
	}

	for _, v := range ll {
		ht.Insert(v)
	}
	fmt.Println("rainier is in the hash table2", ht.Search("rainier"))
	ht.Delete("rainier")
	fmt.Println("rainier is in the hash table2", ht.Search("rainier"))
}