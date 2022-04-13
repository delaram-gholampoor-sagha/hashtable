package main

import "fmt"

// Bucket Structure
// bucket is a linked list in each slot of the hashtable array
type bucket struct {
	// we just need to put the address of the head node in the bucket
	head *BucketNode
}


// Insert
// create a node with the key and insert the node in the bucket.
func (b *bucket) insert(k string) {
	// if k doesnt exist in b in that case i want to execute this
	if !b.Search(k) {
		// we are going to set this new node as a head and then we are going to set
		// this new node, point to the head that was there before.
		newNode := &BucketNode{key: k}
		// next node become the current head
		newNode.next = b.head
		// the head is going to become the new node
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

// Search
// we are going to jump to each of the nodes inside the bucket inside the linked list
// and see if its a match with parameter k
// so we have to put the head in current node
// we want to keep up looping until we find a match
func (b *bucket) Search(k string) bool {
	currentNode := b.head
	// we need to loop until the current loop is empty
	for currentNode != nil {
		// we want to find whether the current node is a match to k
		if currentNode.key == k {
			return true
		}
		// if not we are going to update the node to the next node
		// if we kept on doing this and we went to the end of the node
		// it turns out to be empty
		// we are going to come out of that loop
		currentNode = currentNode.next
	}
	return false

}

// Delete
// delete the node from the bucket
// it will delete the current node by skipping the current node and making the previous node point to the next node
// and we cant do that if we actually go over the previous node
// thats why we are going to put each node into a previous node and express
// the current node as the previous node.next
func (b *bucket) delete(k string) {
	// because we are putting the head in the prevoius node when we are missing one case
	// where the actual matching key is the head
	// if the key of the head is actually a match to what we are trying to delete
	if b.head.key == k {
		// we are reseting the head of this bucket to the second node which is b.head.next
		b.head = b.head.next
		// then we come out of the method
		return
	}
	prevoiusNode := b.head
	for prevoiusNode.next != nil {
		if prevoiusNode.next.key == k {
			// delete
			prevoiusNode.next = prevoiusNode.next.next
			return
		}
		prevoiusNode = prevoiusNode.next
	}
}


