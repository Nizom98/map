package chain_hash_table

import (
	"reflect"
	"testing"
)

func TestItem_Add(t *testing.T) {
	head := &item {
		key: "key1",
		val: []byte{23},
	}
	expectSize := 2

	isAdded := head.set("key2", []byte{23})
	if !isAdded {
		t.Errorf(
			"Got incorrect return value on adding. Got: %v, Expect: %v",
			isAdded,
			true,
		)
		return
	}

	listSize := getListSize(head)
	if listSize != expectSize {
		t.Errorf(
			"Got incorrect list size on adding. Got: %d, Expect: %d",
			listSize,
			expectSize,
		)
	}
}

func TestItem_Update(t *testing.T) {
	head := &item {
		key: "key1",
		val: []byte{23},
	}
	expectSize := 1

	isAdded := head.set("key1", []byte{25})
	if isAdded {
		t.Errorf(
			"Got incorrect return value on updating. Got: %v, Expect: %v",
			isAdded,
			true,
		)
		return
	}

	listSize := getListSize(head)
	if listSize != expectSize {
		t.Errorf(
			"Got incorrect list size on adding. Got: %d, Expect: %d",
			listSize,
			expectSize,
		)
	}
}

func TestItem_Delete(t *testing.T) {
	head := &item {
		key: "key1",
		val: []byte{23},
	}
	expectSize := 0
	isElementDeleted := false

	head, isElementDeleted = head.delete("key1")
	
	if !isElementDeleted {
		t.Errorf(
			"Got incorrect return value on deleting. Got: %v, Expect: %v",
			isElementDeleted,
			true,
		)
		return
	}

	listSize := getListSize(head)
	if listSize != expectSize {
		t.Errorf(
			"Got incorrect list size on deleting. Got: %d, Expect: %d",
			listSize,
			expectSize,
		)
	}
}

func TestItem_Get(t *testing.T) {
	key, value := "key1", "value1"
	expectValue := []byte(value)
	head := &item {
		key: key,
		val: []byte(value),
	}

	gotValue, isExists := head.get("key1")
	if !isExists {
		t.Errorf(
			"Element not found by key(%s) after add. It is incorrect behaviour.",
			key, 
		)
		return
	}

	if !reflect.DeepEqual(expectValue, gotValue) {
		t.Errorf(
			"Got unexpected value on get. Got: %s, Expect: %s",
			string(gotValue),
			string(expectValue),
		)
	}
}

func getListSize(it *item) int {
	if it == nil {
		return 0
	}

	return getListSize(it.next) + 1
}