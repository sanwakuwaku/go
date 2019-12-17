package intset

import (
	"testing"
)

func TestIntSetLen(t *testing.T) {
	s := intSet(1, 34, 123, 3041)
	if s.Len() != 4 {
		t.Errorf("Len() implementation is invalid.")
	}
}

func TestIntSetRemove(t *testing.T) {
	s := intSet(4, 10, 121)

	// 無関係な数値をremoveしても問題ないこと
	s.Remove(1)

	if s.Len() != 3 {
		t.Errorf("Remove() implementation is invalid. irrelevant numbers.")
	}

	if s.Has(1) {
		t.Errorf("Remove() implementation is invalid. ")
	}

	// removeが正しく行えること
	s.Remove(10)
	if s.Len() != 2 {
		t.Errorf("Remove() implementation is invalid. ")
	}

	if s.Has(10) {
		t.Errorf("Remove() implementation is invalid. ")
	}
}

func intSet(values ...int) *IntSet {
	var s IntSet

	for _, value := range values {
		s.Add(value)
	}

	return &s
}
