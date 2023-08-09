package listrank

import "math/rand"

// InitRandomList creates a random permutation of integers from 0 to listSize-1
// and initializes the 'next' slice to represent a random list where next[i]
// is the successor of i in the list.
func InitRandomList(next []int, listSize int) {
	perm := generateSequentialSlice(listSize)
	if listSize < 1 {
		return
	}

	shuffleSlice(perm)

	initializeNextFromPerm(next, perm)
}

// generateSequentialSlice returns a slice of integers from 0 to size-1.
func generateSequentialSlice(size int) []int {
	perm := make([]int, size)
	for i := 0; i < size; i++ {
		perm[i] = i
	}
	return perm
}

// shuffleSlice shuffles the elements of the given slice using the Fisher-Yates algorithm.
func shuffleSlice(slice []int) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// initializeNextFromPerm uses the perm slice to set up the 'next' slice such that
// each element points to its successor based on the perm order. The last element
// in the sequence is set to Nil.
func initializeNextFromPerm(next, perm []int) {
	for i := 0; i < len(perm)-1; i++ {
		next[perm[i]] = perm[i+1]
	}
	next[perm[len(perm)-1]] = Nil
}
