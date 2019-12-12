package main

import (
	"testing"
)

func TestOutline(t *testing.T) {
	if err := outline("http://gopl.io"); err != nil {
		t.Errorf("func outline error:%v\n", err)
	}
}
