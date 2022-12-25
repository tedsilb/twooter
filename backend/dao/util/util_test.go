package util

import (
	"testing"
)

func TestGenerateId(t *testing.T) {
	id1 := GenerateId()
	id2 := GenerateId()

	if len(id1) != 20 {
		t.Errorf("len(id1) want=20, got=%d", len(id1))
	}
	if len(id2) != 20 {
		t.Errorf("len(id2) want=20, got=%d", len(id2))
	}
	if id1 == id2 {
		t.Errorf("id1 == id2, should not be")
	}
}
