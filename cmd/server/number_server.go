package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/jamesrobb/ably-takehome/protocol/generated/protocol"

	"gonum.org/v1/gonum/mathext/prng"
)

const MAX_NUMBERS uint32 = 65535

type numberServer struct {
	protocol.UnimplementedNumbersServer

	clientState     map[uuid.UUID]*State
	clientStateLock sync.Mutex

	stateStorage StateStorage
}

func newNumberServer(stateStore StateStorage) *numberServer {
	return &numberServer{
		stateStorage: stateStore,
	}
}

func (ns *numberServer) GetNumbers(request *protocol.NumbersRequest, stream protocol.Numbers_GetNumbersServer) error {
	var s *State
	var clientID uuid.UUID
	copy(clientID[:], request.ClientId)

	if ns.stateStorage.IsExpiredClientID(clientID) {
		return fmt.Errorf("clientID has expired and cannot be reused")
	}

	// Check if we already have a stored data for a client.
	// Assumptions:
	// - two clients don't ever try to connect with the same clientID.
	// - if a clientID is reused, it is because a client lost connection to the server and is trying to resume.
	storedState, err := ns.stateStorage.GetState(clientID)
	if err == nil {
		s = storedState
		fmt.Printf("found stored state for clientID=%s\n", clientID)
	} else {
		numNumbers := request.NumNumbers
		if numNumbers == 0 {
			return fmt.Errorf("cannot send 0 numbers")
		}
		if numNumbers > MAX_NUMBERS {
			numNumbers = MAX_NUMBERS
		}

		s = &State{
			nextNumber:   0,
			numbersSent:  0,
			totalNumbers: numNumbers,
			lastUpdated:  time.Now(),
			hash:         md5.New(),
			prng:         prng.NewMT19937(),
		}

		// Did the client provide a seed?
		var seed uint64
		if request.Seed > 0 {
			seed = uint64(request.Seed)
		} else {
			seed = uint64(rand.Uint32())
		}
		s.prng.Seed(seed)

		fmt.Printf("sending %d numbers to clientID=%s seed=%d\n", numNumbers, clientID, seed)

		s.nextNumber = s.prng.Uint32()
		io.WriteString(s.hash, fmt.Sprintf("%d", s.nextNumber))
	}

	// Executes loop body every 1 second.
	for range time.Tick(time.Second * 1) {

		isLastPayload := s.totalNumbers == s.numbersSent+1
		var payload *protocol.NumberResponse

		if isLastPayload {
			checksum := hex.EncodeToString(s.hash.Sum(nil)[:])
			payload = &protocol.NumberResponse{
				Number:   s.nextNumber,
				Checksum: checksum,
			}
		} else {
			payload = &protocol.NumberResponse{
				Number:   s.nextNumber,
				Checksum: "",
			}
		}

		err := stream.Send(payload)
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to send number: %s\n", payload)
		}

		if isLastPayload {
			ns.stateStorage.DeleteState(clientID)

			return nil
		}

		// Number successfully sent, can update and store state
		s.nextNumber = s.prng.Uint32()
		s.numbersSent++
		s.lastUpdated = time.Now()
		io.WriteString(s.hash, fmt.Sprintf("%d", s.nextNumber))
		ns.stateStorage.SetState(clientID, s)
	}

	return nil
}
