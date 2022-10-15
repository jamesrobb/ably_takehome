package main

import (
	"fmt"
	"hash"
	"sync"
	"time"

	"github.com/google/uuid"
	"gonum.org/v1/gonum/mathext/prng"
)

const GARBGAGE_TIMEOUT time.Duration = 30 * time.Second

type State struct {
	numbersSent  uint32
	nextNumber   uint32
	totalNumbers uint32
	lastUpdated  time.Time
	hash         hash.Hash
	prng         *prng.MT19937
}

type StateStorage interface {
	IsExpiredClientID(clientID uuid.UUID) bool
	GetState(clientID uuid.UUID) (*State, error)
	SetState(clientID uuid.UUID, state *State) error
	DeleteState(clientID uuid.UUID) error
}

type InMemoryStorage struct {
	states     map[uuid.UUID]*State
	statesLock sync.Mutex

	badClients     map[uuid.UUID]bool
	badClientsLock sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		states:     make(map[uuid.UUID]*State),
		badClients: make(map[uuid.UUID]bool),
	}
}

func (ims *InMemoryStorage) IsExpiredClientID(clientID uuid.UUID) bool {
	// Triggers garbage collection (which marks expired clients)
	ims.statesLock.Lock()
	ims.garbageCollectStates()
	ims.statesLock.Unlock()

	_, ok := ims.badClients[clientID]

	return ok
}

func (ims *InMemoryStorage) GetState(clientID uuid.UUID) (*State, error) {
	ims.statesLock.Lock()
	ims.statesLock.Unlock()

	ims.garbageCollectStates()

	state, ok := ims.states[clientID]
	if !ok {
		return nil, fmt.Errorf("state not found for clientID=%s", clientID)
	}

	return state, nil
}

func (ims *InMemoryStorage) garbageCollectStates() {
	// Assumes calling method holds ims.statesLock.

	for key, state := range ims.states {
		now := time.Now()
		if now.Sub(state.lastUpdated) > GARBGAGE_TIMEOUT {
			ims.badClientsLock.Lock()
			ims.badClients[key] = true
			ims.badClientsLock.Unlock()

			delete(ims.states, key)
		}
	}
}

func (ims *InMemoryStorage) DeleteState(clientID uuid.UUID) error {
	ims.statesLock.Lock()
	defer ims.statesLock.Unlock()

	delete(ims.states, clientID)

	return nil
}

func (ims *InMemoryStorage) SetState(clientID uuid.UUID, state *State) error {
	ims.statesLock.Lock()
	defer ims.statesLock.Unlock()

	storeState := &State{
		numbersSent:  state.numbersSent,
		nextNumber:   state.nextNumber,
		totalNumbers: state.totalNumbers,
		lastUpdated:  state.lastUpdated,
		hash:         state.hash,
		prng:         state.prng,
	}
	ims.states[clientID] = storeState

	return nil
}
