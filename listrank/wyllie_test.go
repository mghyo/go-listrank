package listrank

import (
	"reflect"
	"testing"
)

func TestWyllieListRanks(t *testing.T) {
	tests := []struct {
		listSize int
	}{
		{
			listSize: 1,
		},
		{
			listSize: 10,
		},
		{
			listSize: 100,
		},
		{
			listSize: 1000,
		},
		{
			listSize: 10000,
		},
		{
			listSize: 100000,
		},
	}
	for _, tt := range tests {
		next := make([]int, tt.listSize)
		InitRandomList(next, tt.listSize)

		head := FindHead(next, tt.listSize)

		rankWyllie := make([]int, tt.listSize)
		WyllieListRanks(head, next, rankWyllie, tt.listSize)

		rankSeq := make([]int, tt.listSize)
		SeqListRanks(head, next, rankSeq, tt.listSize)

		if !reflect.DeepEqual(rankWyllie, rankSeq) {
			t.Errorf("For list starting at %d, expected ranks %v but got %v", head, rankSeq, rankWyllie)
		}
	}

}
