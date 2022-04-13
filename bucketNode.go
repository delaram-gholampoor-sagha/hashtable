package main


// BucketNode structure
// is a linked list node that holds the key
// it also needs to have an address pointing to the next node
type BucketNode struct {
	key string
	// the address of another node
	next *BucketNode
}

