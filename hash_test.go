package kv_map

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	mymap := NewWithDefault()

	mymap.Set("key1", []byte("val1"))
	fmt.Println(mymap.Get("key1"))

	fmt.Println(mymap.Get("key2"))
	fmt.Println()

}