package listrank

// SeqListRanks sequentially assigns ranks to each node in the list,
// starting from the head node. The head node gets rank 0, its successor
// gets rank 1, and so on.
func SeqListRanks(head int, next []int, rank []int, listSize int) {
	currentRank := 0
	currentNode := head

	// Traverse the list, assigning ranks sequentially
	for currentNode != Nil {
		rank[currentNode] = currentRank
		currentRank++
		currentNode = next[currentNode]
	}
}
