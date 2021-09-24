package world

import (
	"fmt"
	"testing"
	"github.com/wallyqs/go-mod-a"
)

func TestWorld(t *testing.T) {
	got := World()
	expected := "World"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}

func TestHelloWorld(t *testing.T) {
	got := fmt.Sprintf("%s %s", hello.Hello(), World())
	expected := "Hello World"
	if got != expected {
		t.Fatalf("Expected %v, got: %v", expected, got)
	}
}
