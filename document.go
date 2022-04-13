package main


// What is load factor ?
// The amount of capacity which is to be exhausted for the HashMap to increase its capacity.

// Why load factor ?
// Load factor is by default 0.75 of the initial capacity (16) therefore 25% of 
// the buckets will be free before there is an increase in the capacity & this 
// makes many new buckets with new hashcodes pointing to them to exist just after
//  the increase in the number of buckets.

// Why should you keep many free buckets & what is the impact of keeping free buckets on the performance ?
// If you set the loading factor to say 1.0 then something very interesting might happen.

// Say you are adding an object x to your hashmap whose hashCode is 888 & 
// in your hashmap the bucket representing the hashcode is free , so the 
// object x gets added to the bucket, but now again say if you are adding 
// another object y whose hashCode is also 888 then your object y will get 
// added for sure BUT at the end of the bucket (because the buckets are 
// nothing but linkedList implementation storing key,value & next) now 
// this has a performance impact ! Since your object y is no longer present
// in the head of the bucket if you perform a lookup the time taken is not 
// going to be O(1) this time it depends on how many items are there in the
// same bucket. This is called hash collision by the way & this even happens
// when your loading factor is less than 1.

// Correlation between performance, hash collision & loading factor
// Lower load factor = more free buckets = less chances of collision = high performance = high space requirement.
// Higher load factor = fewer free buckets = higher chance of collision = lower performance = lower space requirement.

// MurmurHash


// Load Factor
// The Load factor is a measure that decides when to increase the HashMap capacity to
// maintain the get() and put() operation complexity of O(1).
// The default load factor of HashMap is 0.75f (75% of the map size).

// how full you want your array to be ?
// loadfactor = n/k
// n = the number of entries that exist 
// k = size 

// The load factor is how many items are in the HashMap as a ratio 
// of the size of the underlying data structure.

// HashMaps are initialised to a certain size. That is if all goes perfect you can
// only put so many key-value pairs in your HashMap before the first collision occurs.
// Collisions slow things down so a new data structure with a higher capacity is created internally.
// The load factor determines when this is done (usually the threshold is 0.75).
