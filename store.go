package main

import (
	"sync"
)

type UVStore struct {
	sync.RWMutex
	seenUsers map[int]struct{}
	histogram map[uint64]*Bucket
}

type Bucket struct {
	UserIds map[int]struct{}
	Count   int
}

func NewBucket() *Bucket {
	return &Bucket{UserIds: make(map[int]struct{})}
}

func (b *Bucket) AddUserId(id int) {
	if _, ok := b.UserIds[id]; !ok {
		b.UserIds[id] = struct{}{}
	}
}

func (b *Bucket) AddCount(count int) {
	b.Count += count
}

func NewUVStore() *UVStore {
	return &UVStore{
		seenUsers: make(map[int]struct{}),
		histogram: make(map[uint64]*Bucket),
	}
}

func (s *UVStore) Add(v Visit) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.histogram[v.Timestamp]; !ok {
		s.histogram[v.Timestamp] = NewBucket()
	}

	s.histogram[v.Timestamp].AddUserId(v.UserId)
	s.histogram[v.Timestamp].AddCount(1)
}

func (s *UVStore) Count(start uint64, end uint64) uint64 {
	s.RLock()
	defer s.RUnlock()

	totalBucket := NewBucket()

	for timestamp, bucket := range s.histogram {
		if timestamp >= start && timestamp <= end {
			for id, _ := range bucket.UserIds {
				totalBucket.AddUserId(id)
			}
			totalBucket.AddCount(bucket.Count)
		}
	}

	return uint64(len(totalBucket.UserIds))
}
