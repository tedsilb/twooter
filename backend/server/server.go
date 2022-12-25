package main

import (
	"context"
	"log"
	"net"

	"github.com/tedsilb/twooter/backend/dao"
	"github.com/tedsilb/twooter/backend/validate"
	pb "github.com/tedsilb/twooter/proto/twooterpb"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement TwooterServer.
type server struct {
	pb.TwooterServer
}

func (s server) CreateTwoot(ctx context.Context, req *pb.CreateTwootRequest) (*pb.Twoot, error) {
	if err := validate.CreateTwootRequest(req); err != nil {
		return nil, err
	}

	t, err := dao.Create(req.GetTwoot())
	if err != nil {
		return nil, err
	}

	log.Printf("Created: %v", t)
	return t, nil
}

func (s server) ListTwoots(ctx context.Context, req *pb.ListTwootsRequest) (*pb.ListTwootsResponse, error) {
	twoots, err := dao.List()
	if err != nil {
		return nil, err
	}

	resp := &pb.ListTwootsResponse{Twoots: *twoots}

	log.Printf("Listed: %v", resp)
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTwooterServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
