package world

import (
	"testing"
)

func TestWorld(t *testing.T) {
	got := World()
	expected := "World"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}
