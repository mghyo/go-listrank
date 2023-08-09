package listrank

import (
	"math"
	"sync"
)

// calculateDepth calculates the required depth based on the list size.
func calculateDepth(listSize int) int {
	return int(math.Ceil(math.Log(float64(listSize)) / math.Log(2)))
}

// initializeAuxiliaryArrays sets up the initial values for rank and auxiliary arrays.
func initializeAuxiliaryArrays(rank, rankTemp, nextTemp1, nextTemp2, next []int, head int) {
	for i := range rank {
		rank[i], rankTemp[i] = 1, 1
		nextTemp1[i], nextTemp2[i] = next[i], next[i]
	}
	rank[head], rankTemp[head] = 0, 0
}

// parallelRankCalculation performs parallel rank calculation using goroutines.
func parallelRankCalculation(rank, rankTemp, nextTemp1 []int, wg *sync.WaitGroup) {
	for j := range rank {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			if nextNode := nextTemp1[j]; nextNode != Nil {
				rankTemp[nextNode] = rank[j] + rank[nextNode]
			}
		}(j)
	}
	wg.Wait()
}

// updateNextPointers updates next pointers in parallel using goroutines.
func updateNextPointers(nextTemp1, nextTemp2 []int, wg *sync.WaitGroup) {
	for j := range nextTemp1 {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			if nextNode := nextTemp1[j]; nextNode != Nil {
				nextTemp2[j] = nextTemp1[nextNode]
			}
		}(j)
	}
	wg.Wait()
}

// WyllieListRanks uses Wyllie's parallel algorithm to assign ranks to the nodes.
// It leverages goroutines for parallel computation, providing efficient rank assignment
// especially for larger lists.
func WyllieListRanks(head int, next []int, rank []int, listSize int) {
	// Initialize auxiliary arrays
	rankTemp := make([]int, listSize)
	nextTemp1 := make([]int, listSize)
	nextTemp2 := make([]int, listSize)

	depth := calculateDepth(listSize)
	initializeAuxiliaryArrays(rank, rankTemp, nextTemp1, nextTemp2, next, head)

	var wg sync.WaitGroup
	// Main computation
	for i := 0; i < depth; i++ {
		parallelRankCalculation(rank, rankTemp, nextTemp1, &wg)
		updateNextPointers(nextTemp1, nextTemp2, &wg)

		// Copy temp arrays back to original arrays
		copy(rank, rankTemp)
		copy(nextTemp1, nextTemp2)
	}
}
