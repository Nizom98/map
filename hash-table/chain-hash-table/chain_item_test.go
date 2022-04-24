package chain_hash_table

import (
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

func getListSize(it *item) int {
	if it == nil {
		return 0
	}

	return getListSize(it.next) + 1
}