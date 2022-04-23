package kv_map

import "sync"

const (
	defaultPolinom       = 4
	defaultHashTableSize = 1000
)

func NewWithDefault() Map {
	return New(defaultPolinom, defaultHashTableSize)
}

func New(polinom, hashTableSize int) Map {
	return &hmap{
		bucket:        make([]*item, hashTableSize),
		count:         0,
		polinom:       polinom,
		hashTableSize: hashTableSize,
		mu:            new(sync.Mutex),
	}
}
