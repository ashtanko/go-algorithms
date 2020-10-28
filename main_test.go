package main

import "testing"

func TestAddition(t *testing.T) {
	if Additional(2, 2) != 4 {
		t.Error("Expected 2 + 2 equals 4")
	}
}
