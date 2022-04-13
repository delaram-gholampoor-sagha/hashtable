package main

// HashTable Structure
// our hash table needs to hold an array
type HashTable struct {
	array [arraySize]*bucket
}


// Insert
// will take in a key and add it to the hashtable array, like "omid" or "ali"
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	// after we get a hash key we turn it into a hashcode and put it into the index
	index := hash(key)
	// we are going to go to that array slot
	// then it will insert the key in that slot
	return h.array[index].Search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}
