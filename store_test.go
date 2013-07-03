package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	//   store := NewUVStore()

	// visits := []Visit{
	//   Visit{UserId: 1, Timestamp: 1},
	//   Visit{UserId: 1, Timestamp: 2},
	//   Visit{UserId: 1, Timestamp: 3},
	// }

	// if store.Add(visits[0]) == false {
	//   t.Errorf("First visit should be added. %v", visits[0])
	// }
	// if bucket, _ := store.histogram[visits[0].Timestamp]; bucket.Count != 1 {
	//   t.Errorf("Value should have been %d. Got: %d", 1, bucket.Count)
	// }

	// for _, visit := range visits[1:] {
	//   if store.Add(visit) == true {
	//     t.Errorf("First visit should not be added. %v", visit)
	//   }
	// }
}

func TestCount(t *testing.T) {
	store := NewUVStore()

	visits := []Visit{
		Visit{UserId: 1, Timestamp: 1},
		Visit{UserId: 1, Timestamp: 2},
		Visit{UserId: 1, Timestamp: 3},
	}

	for _, visit := range visits {
		store.Add(visit)
	}

	for _, visit := range visits {
		start, end, wantCount := uint64(visit.Timestamp), uint64(3), uint64(1)

		if count := store.Count(start, end); count != wantCount {
			t.Errorf("Count failed: want %d, got %d", wantCount, count)
		}
	}
}
