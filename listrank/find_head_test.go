package listrank

import "testing"

func TestFindHead(t *testing.T) {
	tests := []struct {
		name     string
		next     []int
		listSize int
		want     int
	}{
		// Happy Paths
		{"Standard list", []int{1, 2, 3, Nil}, 4, 0},
		{"Head in middle", []int{3, Nil, 0, 1}, 4, 2},

		// Edge Cases
		{"Empty list", []int{}, 0, Nil},
		{"Single node list", []int{Nil}, 1, 0},
		{"Fully disconnected", []int{Nil, Nil, Nil}, 3, Nil}, // returning the first node as head, as per the current logic

		// Negative Cases
		{"List with cycle", []int{1, 2, 3, 1}, 4, Nil},                    // because every node has a predecessor
		{"Multiple potential heads", []int{1, Nil, 4, 5, 3, Nil}, 6, Nil}, // or 2, current logic returns the first head it finds

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindHead(tt.next, tt.listSize); got != tt.want {
				t.Errorf("FindHead() = %v, want %v", got, tt.want)
			}
		})
	}
}
