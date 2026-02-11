package models

import (
	"testing"
	"time"
)

func TestGreetingFields(t *testing.T) {
	now := time.Now()
	g := Greeting{
		ID:        1,
		Message:   "Test message",
		CreatedAt: now,
	}

	if g.ID != 1 {
		t.Errorf("Expected ID 1, got %d", g.ID)
	}
	if g.Message != "Test message" {
		t.Errorf("Expected message 'Test message', got %s", g.Message)
	}
	if !g.CreatedAt.Equal(now) {
		t.Errorf("Expected CreatedAt %v, got %v", now, g.CreatedAt)
	}
}
