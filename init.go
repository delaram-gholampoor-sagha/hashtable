package main


// each slot of the array  has to hold a bucket so we are going to make a function that initilizes this particulare hashtable
// and create a bucket into each slot
// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	// it can loop through each of the slots inside the hash array
	for i := range result.array {
		// and inside there we create a bucket
		result.array[i] = &bucket{}
	}
	return result
}
