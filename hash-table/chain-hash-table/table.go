package chain_hash_table

import (
	"sync"
)

func NewHmap(polinom, hashTableSize uint) *Hmap {
	return &Hmap{
		bucket:        make([]*item, hashTableSize),
		count:         0,
		polinom:       polinom,
		hashTableSize: hashTableSize,
		mu:            new(sync.Mutex),
	}
}

// Set вставка в таблицу значения val по ключу key.
func (hm *Hmap) Set(key string, val []byte) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	position := hash.WithPolinom(key, hm.polinom, hm.hashTableSize)

	if hm.bucket[position] == nil {
		hm.bucket[position] = &item{
			key: key,
			val: val,
		}
		hm.count++
		return
	}

	if addNewElement := hm.bucket[position].set(key, val); addNewElement {
		hm.count++
	}
}

// Get получение из таблицы значения по ключу.
// Если элемент с ключем key отсутствует, то вернется false вторым параметром.
func (hm *Hmap) Get(key string) ([]byte, bool) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	if len(hm.bucket) == 0 {
		return nil, false
	}

	position := hash.WithPolinom(key, hm.polinom, hm.hashTableSize)
	if hm.bucket[position] == nil {
		return nil, false
	}

	return hm.bucket[position].get(key)
}

// Delete удаление элемента по ключу key.
func (hm *Hmap) Delete(key string) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	if len(hm.bucket) == 0 {
		return
	}

	position := hash.WithPolinom(key, hm.polinom, hm.hashTableSize)
	if hm.bucket[position] == nil {
		return
	}

	if updatedList, deleted := hm.bucket[position].delete(key); deleted {
		hm.bucket[position] = updatedList
		hm.count--
	}
}

//Len текущее количество элементов в таблице
func (hm *Hmap) Len() int {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	return hm.count
}