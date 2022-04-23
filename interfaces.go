package kv_map

type Map interface {
	Set(key string, val []byte)
	Get(key string) ([]byte, bool)
	Delete(key string)
	Len() int
}
 