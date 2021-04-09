package storage

import (
	"log"
	"strconv"
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

func GetClients() []int32 {
	mux.Lock()
	defer mux.Unlock()

	ret := make([]int32, 0)
	for k, _ := range cache {
		ret = append(ret, k)
	}
	return ret
}

func GetClient(id string) (string, bool) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf(err.Error())
		return "", false
	}
	mux.Lock()
	defer mux.Unlock()
	client, ok := cache[int32(ID)]
	return client, ok
}
