package main

import (
	"context"
	"testing"

	twootdao "github.com/tedsilb/twooter/backend/dao"
	pb "github.com/tedsilb/twooter/proto/twooterpb"

	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateTwoot(t *testing.T) {
	tests := []struct {
		name        string
		req         *pb.CreateTwootRequest
		expectedErr bool
	}{
		{
			name: "req with normal fields",
			req: &pb.CreateTwootRequest{
				Twoot: &pb.Twoot{
					Creator: &pb.User{
						Id:     "asdf1234",
						Handle: "@someone",
						Email:  "someone@email.com",
					},
					Message: "hello there :)",
				},
			},
			expectedErr: false,
		},
		{
			name: "req with id",
			req: &pb.CreateTwootRequest{
				Twoot: &pb.Twoot{
					Id: "somethingblahblah",
					Creator: &pb.User{
						Id:     "asdf1234",
						Handle: "@someone",
						Email:  "someone@email.com",
					},
					Message: "hello there :)",
				},
			},
			expectedErr: true,
		},
		{
			name: "req with create_time",
			req: &pb.CreateTwootRequest{
				Twoot: &pb.Twoot{
					Creator: &pb.User{
						Id:     "asdf1234",
						Handle: "@someone",
						Email:  "someone@email.com",
					},
					CreateTime: &tspb.Timestamp{Seconds: 123456},
					Message:    "hello there :)",
				},
			},
			expectedErr: true,
		},
		{
			name: "nil twoot",
			req: &pb.CreateTwootRequest{
				Twoot: nil,
			},
			expectedErr: true,
		},
		{
			name:        "nil request",
			req:         nil,
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			n := tspb.Now()
			twooter := &server{}

			got, err := twooter.CreateTwoot(ctx, tc.req)

			if tc.expectedErr {
				if err == nil {
					t.Error("expected error but none returned")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if got.Id == "" {
					t.Errorf("id is missing: %v", got.Id)
				}
				if got.CreateTime == nil {
					t.Errorf("create time is missing: %v", got.CreateTime)
				}
				if got.Message != tc.req.Twoot.Message {
					t.Errorf("message mismatch: got=%s, want=%s", got.Message, tc.req.Twoot.Message)
				}
				if got.Creator.Handle != tc.req.Twoot.Creator.Handle {
					t.Errorf("creator handle mismatch: got=%s, want=%s", got.Creator.Handle, tc.req.Twoot.Creator.Handle)
				}
				if got.Creator.Email != tc.req.Twoot.Creator.Email {
					t.Errorf("creator email mismatch: got=%s, want=%s", got.Creator.Email, tc.req.Twoot.Creator.Email)
				}
				if got.CreateTime.Seconds < n.Seconds && got.CreateTime.Nanos < n.Nanos {
					t.Errorf("create time is lower than expected: %v", got.CreateTime)
				}
			}
		})
	}
}

func TestListTwoots(t *testing.T) {
	tests := []struct {
		name   string
		twoots []*pb.Twoot
	}{
		{
			name: "single twoot",
			twoots: []*pb.Twoot{
				{
					Creator: &pb.User{
						Handle: "@someone",
						Email:  "someone@email.com",
					},
					Message: "hello there :)",
				},
			},
		},
		{
			name: "multiple twoots",
			twoots: []*pb.Twoot{
				{
					Creator: &pb.User{
						Handle: "@someone",
						Email:  "someone@email.com",
					},
					Message: "hello there :)",
				},
				{
					Creator: &pb.User{
						Handle: "@someone",
						Email:  "someone@email.com",
					},
					Message: "hello there again :)",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			twooter := &server{}
			for _, twoot := range tc.twoots {
				_, err := twooter.CreateTwoot(ctx, &pb.CreateTwootRequest{Twoot: twoot})
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

			got, err := twooter.ListTwoots(ctx, &pb.ListTwootsRequest{})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			for i, twoot := range tc.twoots {
				if got.Twoots[i].Message != twoot.Message {
					t.Errorf("message mismatch: got=%s, want=%s", got.Twoots[i].Message, twoot.Message)
				}
				if got.Twoots[i].Creator.Handle != twoot.Creator.Handle {
					t.Errorf("creator handle mismatch: got=%s, want=%s", got.Twoots[i].Creator.Handle, twoot.Creator.Handle)
				}
				if got.Twoots[i].Creator.Email != twoot.Creator.Email {
					t.Errorf("creator email mismatch: got=%s, want=%s", got.Twoots[i].Creator.Email, twoot.Creator.Email)
				}
			}
		})
		twootdao.Clear()
	}
}
