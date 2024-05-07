package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	added := Add(5, 6)

	if added != 11 {
		t.Errorf("Add Calculate Error")
	}
}