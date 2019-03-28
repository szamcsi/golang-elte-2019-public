package lines

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"math"
	"path/filepath"
	"testing"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/texts"
)

const testDir = "../../testdata/"

// OMITTED: func TestLoad(t *testing.T)

type fake struct{} // TODO: use fake for texts_proto.TextsServiceClient

func (f *fake) Count(_ context.Context, req *pb.CountRequest, _ ...grpc.CallOption) (*pb.CountResponse, error) {
	resp := &pb.CountResponse{}
	if len(req.Lines) == 0 {
		return resp, nil
	}
	resp.Min = math.MaxInt32
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

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		path    string
		mmc     *MinMaxCount
		wantErr bool
	}{
		{"9", &MinMaxCount{1, 9, 9}, false},
		{"10", &MinMaxCount{1, 2, 10}, false},
		{"80", &MinMaxCount{0, 80, 9}, false},
		{"12", nil, true},
	} {
		mmc, err := Count(context.Background(), &fake{}, filepath.Join(testDir, tc.path))
		if tc.wantErr != (err != nil) {
			t.Errorf("Count(%q) err != nil is %v; want %v", tc.path, err != nil, tc.wantErr)
			continue
		}
		if !cmp.Equal(mmc, tc.mmc) {
			t.Errorf("Count(%q) = %v, _; want = %v, _", tc.path, mmc, tc.mmc)
		}
	}
}
