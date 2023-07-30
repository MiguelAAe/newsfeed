package rsspuller

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestShort(t *testing.T) {

	fakeTime := time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC)

	items := []Item{
		{
			Title:   "number 1",
			PubDate: fakeTime.Add(1 * time.Hour),
		},
		{
			Title:   "number 0",
			PubDate: fakeTime.Add(-1 * time.Hour),
		},
		{
			Title:   "number 2",
			PubDate: fakeTime.Add(2 * time.Hour),
		},
		{
			Title:   "number -1",
			PubDate: fakeTime.Add(-4 * time.Hour),
		},
	}

	fmt.Println(items)

	sort.Sort(ByItem(items))

	fmt.Println(items)
}
