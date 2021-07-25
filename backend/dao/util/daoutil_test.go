package daoutil

import (
	"testing"
	"time"
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

func TestNow(t *testing.T) {
	before := time.Now().UnixNano()
	got := Now()
	after := time.Now().UnixNano()

	if diff := (got.Seconds + int64(got.Nanos)) - before; diff >= 0 {
		t.Errorf("diff between got and before = %v, should be negative", diff)
	}
	if diff := (after - got.Seconds + int64(got.Nanos)); diff <= 0 {
		t.Errorf("diff between got and after = %v, should be positive", diff)
	}
}
