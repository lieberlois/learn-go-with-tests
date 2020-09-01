package main

import (
	"log"
	"net/http"
	"sync"
)

type InMemoryPlayerStore struct {
	store map[string]int
	mutex sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	handler := &PlayerServer{store: NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}