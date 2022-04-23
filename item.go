package kv_map

type item struct {
	key  string
	val  []byte
	next *item
}

func (it *item) get(key string) ([]byte, bool) {
	if it.key == key {
		return it.val, true
	}

	if it.next == nil {
		return nil, false
	}

	return it.next.get(key)
}

func (it *item) set(key string, val []byte) bool {
	if it.key == key {
		it.val = val
		return false
	}

	if it.next == nil {
		it.next = &item{
			key: key,
			val: val,
		}
		return true
	}

	return it.next.set(key, val)
}

func (it *item) delete(key string) (*item, bool) {
	if it.key == key {
		return it.next, true
	}

	if it.next == nil {
		return it, false
	}

	return it.delete(key)
}
