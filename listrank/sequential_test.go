package listrank

import (
	"reflect"
	"testing"
)

func TestSeqListRanks(t *testing.T) {
	tests := []struct {
		head     int
		next     []int
		listSize int
		expected []int
	}{
		{
			head:     0,
			next:     []int{1, 2, 3, Nil},
			listSize: 4,
			expected: []int{0, 1, 2, 3},
		},
		{
			head:     3,
			next:     []int{2, 4, Nil, 1, 0},
			listSize: 5,
			expected: []int{3, 1, 4, 0, 2},
		},
		{
			head:     1,
			next:     []int{Nil, 0},
			listSize: 2,
			expected: []int{1, 0},
		},
		{
			head:     0,
			next:     []int{Nil},
			listSize: 1,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		rank := make([]int, tt.listSize)
		SeqListRanks(tt.head, tt.next, rank, tt.listSize)

		if !reflect.DeepEqual(rank, tt.expected) {
			t.Errorf("For list starting at %d, expected ranks %v but got %v", tt.head, tt.expected, rank)
		}
	}
}
