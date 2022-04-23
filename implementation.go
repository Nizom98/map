package kv_map

import "sync"

type hmap struct {
	bucket        []*item
	count         int
	
	polinom       int
	hashTableSize int

	mu *sync.Mutex
}

func (hm *hmap) Set(key string, val []byte) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	index := hash(key, hm.polinom, hm.hashTableSize)

	if hm.bucket[index] == nil {
		hm.bucket[index] = &item{
			key: key,
			val: val,
		}
		hm.count++
		return
	}

	if addNewElement := hm.bucket[index].set(key, val); addNewElement {
		hm.count++
	}
}

func (hm *hmap) Get(key string) ([]byte, bool) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	if len(hm.bucket) == 0 {
		return nil, false
	}

	index := hash(key, hm.polinom, hm.hashTableSize)
	if hm.bucket[index] == nil {
		return nil, false
	}

	return hm.bucket[index].get(key)
}

func (hm *hmap) Delete(key string) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	if len(hm.bucket) == 0 {
		return
	}

	index := hash(key, hm.polinom, hm.hashTableSize)
	if hm.bucket[index] == nil {
		return
	}

	if updatedList, deleted := hm.bucket[index].delete(key); deleted {
		hm.bucket[index] = updatedList
		hm.count--
	}
}

func (hm *hmap) Len() int {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	return hm.count
}