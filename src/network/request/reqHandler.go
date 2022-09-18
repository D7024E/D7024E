package requestHandler

import (
	"errors"
	"sync"

	"D7024E/log"
)

type request struct {
	id      string
	reponse []byte
}

type requestTable struct {
	reqTable []request
}

var instance *requestTable
var lock = &sync.Mutex{}

func GetInstance() *requestTable {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.INFO("Creating new request table")
			instance = &requestTable{}
		}
	}
	return instance
}

func (rt *requestTable) NewRequest(reqID string) error {
	lock.Lock()
	defer lock.Unlock()
	var newReq request
	newReq.id = reqID

	for i := 0; i < len(rt.reqTable); i++ {
		if rt.reqTable[i].id == reqID {
			return errors.New("Table conflict, request id already in use")
		}
	}

	rt.reqTable = append(rt.reqTable, newReq)
	return nil
}

// Writes a received RPC response to the corresponding request id.
func (rt *requestTable) WriteRespone(reqID string, message []byte) error {
	lock.Lock()
	defer lock.Unlock()

	for i := 0; i < len(rt.reqTable); i++ {
		if rt.reqTable[i].id == reqID {
			if rt.reqTable[i].reponse != nil {
				return errors.New("Table conflict, message already received")
			} else {
				rt.reqTable[i].reponse = message
				return nil
			}
		}
	}
	return errors.New("Response error, no request id match")
}

// Reads the response message to the given request id and clears the request from the table.
func (rt *requestTable) ReadResponse(reqID string, res *[]byte) error {
	lock.Lock()
	defer lock.Unlock()

	for i := 0; i < len(rt.reqTable); i++ {
		if rt.reqTable[i].id == reqID {
			if rt.reqTable[i].reponse != nil {
				*res = rt.reqTable[i].reponse
				// Set the element at index "i" to the last element of the slice and then cut out the last element.
				// This effectively deletes the element at index "i", since the order is not important this is faster than appending each
				// element of the slice after index i to the slice before index "i".
				rt.reqTable[i] = rt.reqTable[len(rt.reqTable)-1]
				rt.reqTable = rt.reqTable[:len(rt.reqTable)-1]
				return nil
			} else {
				rt.reqTable[i] = rt.reqTable[len(rt.reqTable)-1]
				rt.reqTable = rt.reqTable[:len(rt.reqTable)-1]
				return errors.New("Response warning, no response received")
			}
		}
	}

	return errors.New("Response error, no request id match")
}
