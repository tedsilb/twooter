package twootdao

import (
	"fmt"

	util "github.com/tedsilb/twooter/backend/dao/util"
	pb "github.com/tedsilb/twooter/proto/twooterpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	twoots []*pb.Twoot
)

// Clear deletes all Twoots. Only used for tests.
func Clear() error {
	twoots = nil
	return nil
}

// Create creates a Twoot.
func Create(t *pb.Twoot) (*pb.Twoot, error) {
	t.Id = util.GenerateId()
	t.CreateTime = util.Now()

	twoots = append(twoots, t)

	return Get(t.Id)
}

// Get retrieves a Twoot by ID.
func Get(id string) (*pb.Twoot, error) {
	for _, twoot := range twoots {
		if twoot.Id == id {
			return twoot, nil
		}
	}
	return nil, status.Error(codes.NotFound, fmt.Sprintf("twoots/%s not found", id))
}

// List returns all Twoots.
func List() (*[]*pb.Twoot, error) {
	return &twoots, nil
}
