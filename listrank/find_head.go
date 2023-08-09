package listrank

// SeqFindHead identifies the head of the list by finding the node
// that does not have a predecessor. It returns the index of the head node.
func FindHead(next []int, listSize int) int {
	if hasCycle(next) {
		return Nil
	}

	hasPredecessor := initializePredecessorArray(listSize)

	markPredecessors(next, hasPredecessor)

	numHeads := countHeads(hasPredecessor)
	if numHeads > 1 {
		return Nil
	}

	return findHeadIndex(hasPredecessor)
}

// Additional function to handle cycle detection, if required for the tests.
// Detects a cycle in a list and returns true if there is one.
func hasCycle(next []int) bool {
	if len(next) < 2 {
		return false
	}

	slow, fast := 0, next[0]

	for fast != Nil && next[fast] != Nil {
		if slow == fast {
			return true
		}
		slow = next[slow]
		fast = next[next[fast]]
	}

	return false
}

// initializePredecessorArray creates and initializes an array to track
// nodes that have predecessors. All nodes are initialized to have no predecessor (Nil).
func initializePredecessorArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = Nil
	}
	return arr
}

// markPredecessors goes through the 'next' slice and marks nodes
// that are pointed to by others (i.e., they have a predecessor).
func markPredecessors(next, hasPredecessor []int) {
	for _, nextNode := range next {
		if nextNode != Nil {
			hasPredecessor[nextNode] = 1 // marking the node as having a predecessor
		}
	}
}

// countHeads iterates through the hasPredecessor slice to determine how many nodes
// do not have predecessors, indicating they are head nodes.
//
// Parameters:
//   - hasPredecessor: A slice where each entry indicates whether the corresponding
//                     node in a linked list has a predecessor. An entry with a value
//                     of Nil indicates that the node does not have a predecessor.
//
// Returns:
//   - int: The number of head nodes found in the hasPredecessor slice.

func countHeads(hasPredecessor []int) int {
	headCount := 0
	for _, pred := range hasPredecessor {
		if pred == Nil {
			headCount++
		}
	}
	return headCount
}

// findHeadIndex identifies the first node that does not have a predecessor
// and returns its index. If no such node is found, it returns Nil.
func findHeadIndex(hasPredecessor []int) int {
	for i, hasPred := range hasPredecessor {
		if hasPred == Nil {
			return i
		}
	}
	return Nil
}
