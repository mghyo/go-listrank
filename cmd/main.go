package main

import (
	"flag"
	"fmt"
	"github.com/mghyo/go-listrank/listrank"
	"strings"
	"time"
)

var runSeqListRanks, runWyllieListRanks, runHbjListRanks bool

func verifyCorrectness(next []int, rankSeq []int, rankParallel []int, listSize int) bool {
	for i := 0; i < listSize; i++ {
		if rankSeq[i] != rankParallel[i] {
			return false
		}
	}
	return true
}

func formatDuration(d time.Duration) string {
	h := d.Hours()
	m := d.Minutes()
	s := d.Seconds()
	ms := d.Milliseconds()
	us := d.Microseconds()
	ns := d.Nanoseconds()

	switch {
	case h >= 1:
		return fmt.Sprintf("%.2fh", h)
	case m >= 1:
		return fmt.Sprintf("%.2fm", m)
	case s >= 1:
		return fmt.Sprintf("%.2fs", s)
	case ms >= 1:
		return fmt.Sprintf("%dms", ms)
	case us >= 1:
		return fmt.Sprintf("%dµs", us)
	default:
		return fmt.Sprintf("%dns", ns)
	}
}

func testListSize(listSize int) {
	fmt.Printf("%-10d", listSize)

	next := make([]int, listSize)
	listrank.InitRandomList(next, listSize)
	head := listrank.FindHead(next, listSize)

	var rankSeq []int
	if runSeqListRanks {
		rankSeq = make([]int, listSize)
		startTime := time.Now()
		listrank.SeqListRanks(head, next, rankSeq, listSize)
		fmt.Printf("| %-25s", formatDuration(time.Since(startTime)))
	}

	if runWyllieListRanks {
		rankWyllie := make([]int, listSize)
		startTime := time.Now()
		listrank.WyllieListRanks(head, next, rankWyllie, listSize)
		fmt.Printf("| %-25s", formatDuration(time.Since(startTime)))
		if verifyCorrectness(next, rankSeq, rankWyllie, listSize) {
			fmt.Printf("| %-30s", "Correct")
		} else {
			fmt.Printf("| %-30s", "Incorrect")
		}
	}

	if runHbjListRanks {
		rankHbj := make([]int, listSize)
		startTime := time.Now()
		listrank.HbjListRanks(head, next, rankHbj, listSize)
		fmt.Printf("| %-25s", formatDuration(time.Since(startTime)))
		if verifyCorrectness(next, rankSeq, rankHbj, listSize) {
			fmt.Printf("| %-30s", "Correct")
		} else {
			fmt.Printf("| %-30s", "Incorrect")
		}
	}

	fmt.Println()
}

func main() {
	flag.BoolVar(&runSeqListRanks, "seq", true, "Run SeqListRanks algorithm")
	flag.BoolVar(&runWyllieListRanks, "wyllie", true, "Run WyllieListRanks algorithm")
	flag.BoolVar(&runHbjListRanks, "hbj", true, "Run HbjListRanks algorithm")
	flag.Parse()

	sizes := []int{1, 20, 100, 500, 1000, 10000, 100000, 500000}

	var headerColumns []string
	headerColumns = append(headerColumns, fmt.Sprintf("%-10s", "List Size"))
	if runSeqListRanks {
		headerColumns = append(headerColumns, fmt.Sprintf("| %-25s", "SeqListRanks Time"))
	}
	if runWyllieListRanks {
		headerColumns = append(headerColumns, fmt.Sprintf("| %-25s| %-30s", "WyllieListRanks Time", "WyllieListRanks Validation"))
	}
	if runHbjListRanks {
		headerColumns = append(headerColumns, fmt.Sprintf("| %-25s| %-30s", "HbjListRanks Time", "HbjListRanks Validation"))
	}

	header := strings.Join(headerColumns, "")
	headerLine := strings.Repeat("-", len(header))
	fmt.Println(headerLine)
	fmt.Println(header)
	fmt.Println(headerLine)

	for _, size := range sizes {
		testListSize(size)
	}
}
