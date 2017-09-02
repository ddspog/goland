// Package collections contains utility functions for creating specified collections.
package collections

// Set it's a collection of objects that will have expertise in finding keys. Therefore, it's a map that stores 0 bytes
// structs. This way, someone can insert a new key, and remove a key and find if a key it's in the Set.
type Set map[string]struct{}

var exists = struct{}{}

// NewSet creates a new set inserting the values received. It creates a Set, a map storing 0 bytes values, to help on
// the sole purpose of finding the keys inserted.
func NewSet(values ...string) Set {
	var newSet = make(map[string]struct{})

	for _, value := range values {
		newSet[value] = exists
	}
	return newSet
}
