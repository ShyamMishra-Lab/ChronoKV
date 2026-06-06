package store

import (
	"testing"
)

func TestSet (t *testing.T) {
	s := NewStore()
	
	s.Set("Name", "Shyam")

	val, err := s.Get("Name")
	if err != nil {
		t.Errorf("Expected to find key 'Name', but got error: %v", err)
	}
	if val != "Shyam" {
		t.Errorf("Expected value 'Shyam', but got '%s'", val)
	}
}

func TestGet (t *testing.T) {
	s := NewStore()
	
	val, err := s.Get("NonExistentKey")
	if err == nil {
		t.Errorf("Expected error for non-existent key, but got value: %s", val)
	}
}

func TestDelete (t *testing.T) {
	s := NewStore()

	s.Set("Name", "Shyam")
	s.Delete("Name")

	val, err := s.Get("Name")
	if err == nil {
		t.Errorf("Expected error for deleted key, but got value: %s", val)
	}
}