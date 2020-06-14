package main

import "testing"


func TestNewUF(t *testing.T) {
	uf := NewUF(5)
	if uf.count != 5 {
		t.Fatal("count should be 5")
	}
}