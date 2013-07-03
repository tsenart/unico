package main

import (
	"log"
)

type Worker struct {
	id     int
	visits <-chan Visit
	store  *UVStore
}

func NewWorker(id int, visits <-chan Visit, store *UVStore) *Worker {
	return &Worker{id: id, visits: visits, store: store}
}

func (w *Worker) Start() {
	log.Printf("Worker %d: started", w.id)
	go w.loop()
}

func (w *Worker) loop() {
	for visit := range w.visits {
		log.Printf("Worker %d: got visit: %v", w.id, visit)
		w.store.Add(visit)
		log.Printf("Worker %d: stored: %v", w.id, visit)
	}
}
