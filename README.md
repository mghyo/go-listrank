# go-listrank
go-listrank is a Go package that provides algorithms to compute ranks for nodes in a linked list. It includes three algorithms: SeqListRanks, WyllieListRanks, and HbjListRanks.

## Installation
To use go-listrank, you need to have Go installed on your system. You can install the package using the following command:

```bash
go get github.com/mghyo/go-listrank
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/mghyo/go-listrank/listrank"
)

func main() {
	// Example usage of SeqListRanks
	next := []int{1, 2, 3, 0} // Example list, replace this with your linked list
	head := listrank.FindHead(next, len(next))
	rank := make([]int, len(next))
	listrank.SeqListRanks(head, next, rank, len(next))

	// Print the ranks
	for i, r := range rank {
		fmt.Printf("Node %d has rank %d\n", i, r)
	}
}
```
## Problem
The list ranking problem involves determining the position, or rank, of each item in a linked list. That is, the first item in the list should be assigned the number 1, the second item in the list should be assigned the number 2, etc.

## Algorithms

### SeqListRanks
SeqListRanks is a sequential algorithm to assign ranks to nodes in a linked list. It identifies the head of the list, and then sequentially traverses the list to compute the rank for each node.

### WyllieListRanks
WyllieListRanks is an algorithm that uses parallel processing to compute ranks for nodes in a linked list. It divides the list into sublists and processes each sublist in parallel to compute ranks. This algorithm is more efficient when the number of processors is ~ list size


### HbjListRanks
HbjListRanks is another parallel algorithm to assign ranks to nodes in a linked list. It segments the list into sublists based on the number of CPUs and log of the list size, and then processes each sublist in parallel to compute ranks. It is the most efficient algorithm of the three for large linked lists.

### Source
[List Ranking Wiki](https://en.wikipedia.org/wiki/List_ranking)

[List Ranking Description](https://en.algorithmica.org/hpc/external-memory/list-ranking/)

[Prefix Computations on Symmetric Multiprocessors](https://api.drum.lib.umd.edu/server/api/core/bitstreams/3611ff6f-257e-4e34-bbe1-516cdeda00af/content)

## License
This project is licensed under the MIT License. See the LICENSE file for details.