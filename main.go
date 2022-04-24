package kv_map

import (
	cht "pets/map/hash-table/chain-hash-table"
)

const (
	defaultPolinom       = 4
	defaultHashTableSize = 1000
)

func NewWithDefault() Map {
	return New(defaultPolinom, defaultHashTableSize)
}

func New(polinom, hashTableSize int) Map {
	return cht.NewHmap(defaultPolinom, defaultHashTableSize)
}
