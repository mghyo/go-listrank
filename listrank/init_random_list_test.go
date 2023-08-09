package listrank

import (
	"reflect"
	"testing"
)

func TestGenerateSequentialSlice(t *testing.T) {
	tests := []struct {
		size     int
		expected []int
	}{
		{5, []int{0, 1, 2, 3, 4}},
		{0, []int{}},
		{1, []int{0}},
		{3, []int{0, 1, 2}},
		{10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := generateSequentialSlice(tt.size)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Fatalf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestInitializeNextFromPerm(t *testing.T) {
	tests := []struct {
		perm     []int
		expected []int
	}{
		{[]int{4, 2, 1, 0, 3}, []int{3, 0, 1, Nil, 2}},
		{[]int{0}, []int{Nil}},
		{[]int{0, 1}, []int{1, Nil}},
		{[]int{1, 0, 3, 2}, []int{3, 0, Nil, 2}},
		{[]int{4, 0, 3, 1, 2}, []int{3, 2, Nil, 1, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			next := make([]int, len(tt.perm))
			initializeNextFromPerm(next, tt.perm)
			if !reflect.DeepEqual(next, tt.expected) {
				t.Fatalf("Expected %v, but got %v", tt.expected, next)
			}
		})
	}
}

func TestInitRandomList(t *testing.T) {
	tests := []struct {
		listSize int
	}{
		{10},
		{0},
		{1},
		{3},
		{20},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			next := make([]int, tt.listSize)
			InitRandomList(next, tt.listSize)

			// Ensure that all values are unique and within bounds.
			seen := make(map[int]bool)
			for _, val := range next {
				if val == Nil {
					continue
				}
				if val < 0 || val >= tt.listSize {
					t.Fatalf("Value %d out of bounds", val)
				}
				if seen[val] {
					t.Fatalf("Duplicate value: %d", val)
				}
				seen[val] = true
			}
		})
	}
}

func TestShuffleSlice(t *testing.T) {
	tests := []struct {
		original []int
	}{
		{[]int{0, 1, 2, 3, 4}},
		{[]int{}},
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{0, 1, 2}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			shuffled := make([]int, len(tt.original))
			copy(shuffled, tt.original)
			shuffleSlice(shuffled)

			// Check if the shuffled slice has the same elements as the original.
			originalMap := make(map[int]int)
			shuffledMap := make(map[int]int)

			for _, v := range tt.original {
				originalMap[v]++
			}
			for _, v := range shuffled {
				shuffledMap[v]++
			}

			if !reflect.DeepEqual(originalMap, shuffledMap) {
				t.Fatalf("Shuffled slice does not contain the same elements as the original")
			}
		})
	}
}
