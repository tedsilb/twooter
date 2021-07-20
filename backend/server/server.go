package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/tedsilb/twooter/proto/twooterpb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	twoots      []*pb.Twoot
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateId() string {
	n := 20
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// server is used to implement TwooterServer.
type server struct {
	pb.TwooterServer
}

func (s *server) CreateTwoot(ctx context.Context, req *pb.CreateTwootRequest) (*pb.Twoot, error) {
	t := req.GetTwoot()
	log.Printf("Creating: %v", t)
	t.Id = generateId()
	twoots = append(twoots, t)
	log.Printf("Created: %v", t)
	return t, nil
}

func (s *server) ListTwoots(ctx context.Context, req *pb.ListTwootsRequest) (*pb.ListTwootsResponse, error) {
	resp := &pb.ListTwootsResponse{}
	resp.Twoots = twoots
	log.Printf("Listing: %v", resp)
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
