package chain_hash_table

import (
	"reflect"
	"testing"
)

const (
	defaultHashTableSize = 1000
	defaultPolinom       = 4
)

func TestHmap_GetAndSet(t *testing.T) {
	myMap := NewHmap(defaultPolinom, defaultHashTableSize)
	key, value := "key1", "value1"

	_, isExists := myMap.Get(key)
	if isExists {
		t.Errorf(
			"Unexpected exists by key %s", key,
		)
		return
	}

	myMap.Set(key, []byte(value))
	gotValue, isExists := myMap.Get(key)
	if !isExists {
		t.Errorf("Unexpected not exists mark")
		return
	}

	if !reflect.DeepEqual(gotValue, []byte(value)) {
		t.Errorf(
			"Incorrect got value. Got: %s, Expect: %s",
			string(gotValue),
			value,
		)
		return
	}
}

func TestHmap_Len(t *testing.T) {
	myMap := NewHmap(defaultPolinom, defaultHashTableSize)
	lenBefore := myMap.Len()
	
	if lenBefore != 0 {
		t.Errorf(
			"Expect empty hash table, but got %d", lenBefore,
		)
		return
	}

	myMap.Set("key1", []byte("value1"))
	lenAfterSet := myMap.Len()
	if lenAfterSet != 1 {
		t.Errorf("Expect length 1, but got %d", lenAfterSet)
	}

	myMap.Delete("key1")
	lenAfterDelete := myMap.Len()
	if lenAfterDelete != 0 {
		t.Errorf("Expect empty hash table after delete, but got %d", lenAfterDelete)
	}
}