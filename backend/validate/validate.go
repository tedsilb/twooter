package validate

import (
	"fmt"

	pb "github.com/tedsilb/twooter/proto/twooterpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateTwootRequest(req *pb.CreateTwootRequest) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "Request message must be provided")
	}
	if req.GetTwoot() == nil {
		return status.Error(codes.InvalidArgument, "Twoot must be provided")
	}
	if req.GetTwoot().Id != "" {
		return status.Error(codes.InvalidArgument, "ID must not be set, it will be determined once the Twoot is created.")
	}
	fmt.Printf("%v\n", req.GetTwoot().CreateTime)
	if req.GetTwoot().CreateTime != nil {
		return status.Error(codes.InvalidArgument, "Create time must not be set, it will be determined once the Twoot is created.")
	}
	return nil
}
