// Package linesservice deals with lines in files.
package linesservice

//go:generate protoc -I../../proto/lines --go_out=plugins=grpc:../../proto/lines lines.proto
import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"math"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/lines"
)

type service struct{}

// New creates a lines service.
func New() pb.LinesServiceServer {
	return &service{}
}

func (s *service) Count(_ context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	if len(req.Lines) == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument,
			"must provide at least one line in the request")
	}
	resp := &pb.CountResponse{Min: math.MaxInt32}
	for _, l := range req.Lines {
		resp.Count++
		if int32(len(l)) < resp.Min {
			resp.Min = int32(len(l))
		}
		if int32(len(l)) > resp.Max {
			resp.Max = int32(len(l))
		}
	}
	return resp, nil
}
