package main

import (
	"sync"
)

type UVStore struct {
	sync.RWMutex
	histogram map[uint64]Set
}

type Set map[int]struct{}

func (set Set) Add(id int) {
	if _, ok := set[id]; !ok {
		set[id] = struct{}{}
	}
}

func NewUVStore() *UVStore {
	return &UVStore{
		histogram: make(map[uint64]Set),
	}
}

func (s *UVStore) Add(v Visit) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.histogram[v.Timestamp]; !ok {
		s.histogram[v.Timestamp] = make(Set)
	}
	s.histogram[v.Timestamp].Add(v.UserId)
}

func (s *UVStore) Count(start uint64, end uint64) uint64 {
	s.RLock()
	defer s.RUnlock()

	rangeSet := make(Set)
	for timestamp, set := range s.histogram {
		if timestamp >= start && timestamp <= end {
			for id, _ := range set {
				rangeSet.Add(id)
			}
		}
	}
	return uint64(len(rangeSet))
}
