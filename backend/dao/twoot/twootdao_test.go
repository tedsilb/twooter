package twootdao

import (
	pb "github.com/tedsilb/twooter/proto/twooterpb"
	"google.golang.org/protobuf/proto"
	"testing"
)

var (
	twoot = &pb.Twoot{
		Creator: &pb.User{
			Id:     "asdf1234",
			Handle: "@someone",
			Email:  "someone@email.com",
		},
		Message: "hello there :)",
	}
)

func TestCreate(t *testing.T) {
	got, err := Create(twoot)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got.Id == "" {
		t.Errorf("id is missing: %v", got.Id)
	}
	if got.CreateTime == nil {
		t.Errorf("create time is missing: %v", got.CreateTime)
	}
	if got.Message != twoot.Message {
		t.Errorf("message mismatch: got=%s, want=%s", got.Message, twoot.Message)
	}
	if !proto.Equal(got.Creator, twoot.Creator) {
		t.Errorf("creator mismatch: got=%v, want=%v", got.Creator, twoot.Creator)
	}
}

func TestGet(t *testing.T) {
	created, err := Create(twoot)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got, err := Get(created.Id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !proto.Equal(got, twoot) {
		t.Errorf("twoot mismatch: got=%v, want=%v", got, twoot)
	}
}

func TestList(t *testing.T) {
	if _, err := Create(twoot); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := Create(twoot); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got, err := List()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for _, gotTwoot := range *got {
		if !proto.Equal(gotTwoot, twoot) {
			t.Errorf("twoots mismatch: got=%v, want=%v", gotTwoot, twoot)
		}
	}
}
