package storage

import (
	"sync"
	"sync/atomic"
)

var mux sync.Mutex
var cache map[int32]string
var id int32

func init() {
	cache = make(map[int32]string, 0)
}

func getNextID() int32 {
	atomic.AddInt32(&id, 1)
	return id
}

func SetClient(host string) {
	mux.Lock()
	defer mux.Unlock()
	cache[getNextID()] = host
}

func GetClients() map[int32]string {
	mux.Lock()
	defer mux.Unlock()
	return cache
}
