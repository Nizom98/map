package hash

import (
	"testing"
)

func TestWithPolinom_EmptyInputValue(t *testing.T) {
	polinom, hts := uint(34), uint(15000)
	inputValue := ""
	expectedResult := uint(0)
	result := WithPolinom(inputValue, polinom, hts)

	if result != expectedResult {
		t.Errorf(
			"Unexpected result on empty input value. Got: %d, Expect: %d", 
			result, expectedResult,
		)
	}
}

func TestWithPolinom_ZeroHashTableSize(t *testing.T) {
	polinom, hts := uint(34), uint(0)
	inputValue := "dsvgsdg"
	expectedResult := uint(0)
	result := WithPolinom(inputValue, polinom, hts)

	if result != expectedResult {
		t.Errorf(
			"Unexpected result on zero hash table size. Got: %d, Expect: %d", 
			result, expectedResult,
		)
	}
}