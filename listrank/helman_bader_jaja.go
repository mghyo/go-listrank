package listrank

import (
	"math"
	"runtime"
	"sync"
)

// isSubHead checks whether a given node is a head of a sublist or not.
// It returns true if the node is a sub-head, otherwise false.
func isSubHead(node int, head int, subListSize int) bool {
	if node == Nil {
		return false
	}
	return head == node || (head >= subListSize && node < (subListSize-1)) || (head < subListSize && node < subListSize)
}

// linkSingleNodeInSublist links a single node within a sublist to its corresponding position
// in the main list and computes the head rank of the sublist.
//
// Parameters:
// - i: the index of the sublist.
// - subHead: slice containing starting nodes of sublists.
// - nextHead: slice representing the 'next' pointers of the sublist heads.
// - next: slice representing the 'next' pointers in the list.
// - rank: slice containing the rank of each node.
// - head: the head node of the entire list.
// - subListSize: the number of nodes in each sublist.
// - wg: pointer to a WaitGroup used to synchronize goroutines.
func linkSingleNodeInSublist(i int, subHead, nextHead, next, rank []int, head, subListSize int, wg *sync.WaitGroup) {
	defer wg.Done()
	node := subHead[i]
	r := 0
	for node != Nil {
		r++
		node = next[node]
		if isSubHead(node, head, subListSize) {
			if head >= subListSize && i == subListSize-1 {
				nextHead[subListSize-1] = node
			} else {
				nextHead[subHead[i]] = node
			}
			rank[node] = r
			break
		}
	}
}

// linkSublists uses goroutines to link the sublists and compute their head ranks.
// It uses synchronization with a WaitGroup to ensure all goroutines complete.
func linkSublists(subHead, nextHead, next, rank []int, head, subListSize int, wg *sync.WaitGroup) {
	for i := 0; i < subListSize; i++ {
		wg.Add(1)
		go linkSingleNodeInSublist(i, subHead, nextHead, next, rank, head, subListSize, wg)
	}
	wg.Wait()
}

// rankSubheads assigns ranks to heads of sublists.
// The function traverses the list and assigns ranks based on the head nodes.
func rankSubheads(subHead, nextHead, rank []int, head, subListSize int) {
	node := head
	r := 0
	for node != Nil {
		if head == node {
			if head >= subListSize {
				node = nextHead[subListSize-1]
			} else {
				node = nextHead[node]
			}
		} else {
			rank[node] += r
			r = rank[node]
			node = nextHead[node]
		}
	}
}

// rankNodeInSublist ranks nodes within a single sublist.
// The function traverses the sublist and assigns ranks to its nodes.
// It is designed to be run as a goroutine and works in parallel with other instances.
//
// Parameters:
// - i: the index of the sublist.
// - subHead: slice containing starting nodes of sublists.
// - next: slice representing the 'next' pointers in the list.
// - rank: slice containing the rank of each node.
// - head: the head node of the entire list.
// - subListSize: the number of nodes in each sublist.
// - wg: pointer to a WaitGroup used to synchronize goroutines.
func rankNodeInSublist(i int, subHead, next, rank []int, head, subListSize int, wg *sync.WaitGroup) {
	defer wg.Done()
	node := subHead[i]
	r := rank[node]
	for node != Nil {
		rank[node] = r
		r++
		node = next[node]
		if isSubHead(node, head, subListSize) {
			break
		}
	}
}

// rankSublists uses goroutines to rank nodes in sublists in parallel.
// It traverses each sublist and assigns ranks to its nodes.
func rankSublists(subHead, next, rank []int, head, subListSize int, wg *sync.WaitGroup) {
	for i := 0; i < subListSize; i++ {
		wg.Add(1)
		go rankNodeInSublist(i, subHead, next, rank, head, subListSize, wg)
	}
	wg.Wait()
}

// calculateSubListSize computes the size of each sublist based on the number of CPUs and the logarithm of the list size.
// It ensures the sublist size is never larger than the list size itself and never less than 1.
func calculateSubListSize(listSize int) int {
	numCpus := runtime.NumCPU()
	subListSize := int(float64(numCpus) * math.Ceil(math.Log(float64(listSize))/math.Log(2)))
	subListSize = int(math.Min(float64(subListSize), float64(listSize)))
	subListSize = int(math.Max(float64(subListSize), 1))
	return subListSize
}

// HbjListRanks uses the parallel algorithm by Hellman, Bader, and Jaja to assign ranks to nodes.
// It segments the list into sublists based on the number of CPUs and log of the list size,
// and then processes each sublist in parallel.
func HbjListRanks(head int, next, rank []int, listSize int) {
	subListSize := calculateSubListSize(listSize)

	subHead := make([]int, subListSize)
	nextHead := make([]int, subListSize)
	for i := range subHead {
		subHead[i] = i
		nextHead[i] = Nil
	}
	if head >= subListSize {
		subHead[subListSize-1] = head
	}
	rank[head] = 0

	var wg sync.WaitGroup

	linkSublists(subHead, nextHead, next, rank, head, subListSize, &wg)
	rankSubheads(subHead, nextHead, rank, head, subListSize)
	rankSublists(subHead, next, rank, head, subListSize, &wg)
}
