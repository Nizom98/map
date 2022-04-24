package chain_hash_table

// get получаем значение из списка по ключу.
// Если элемент с таким ключем отсутствует, то вернется false вторым параметром.
func (it *item) get(key string) ([]byte, bool) {
	if it.key == key {
		return it.val, true
	}

	if it.next == nil {
		return nil, false
	}

	return it.next.get(key)
}

// set вставка в список элемента по ключу.
// Если элемент с таким ключем уже существует, 
// то занчение существующего ключа будет обновлена и вернется false.
// Если элемент с таким ключем отсутствует, то этот элемент будет добавлен и вернется true.
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

// delete удаление из списка элемента по ключу.
// Первый возвращаемый параметр - указатель на обновленный список.
// Второй возвращаемый параметр - true если произошло удаление, 
// иначе false(элемент с ключем key не найден).
func (it *item) delete(key string) (*item, bool) {
	if it.key == key {
		return it.next, true
	}

	if it.next == nil {
		return it, false
	}

	return it.delete(key)
}
