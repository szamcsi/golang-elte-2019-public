package lines

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"path/filepath"
	"testing"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/lines"
)

const testDir = "../../testdata/"

// OMITTED: func TestLoad(t *testing.T)

type fake struct{}

func (f *fake) Count(_ context.Context, req *pb.CountRequest, _ ...grpc.CallOption) (*pb.CountResponse, error) {
	// TODO: implement a stub that will always return Min 1, Max 2, Count 3.
}

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		path    string
		mmc     *MinMaxCount
		wantErr bool
	}{
		{"9", &MinMaxCount{1, 2, 3}, false},
		{"10", &MinMaxCount{1, 2, 3}, false},
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
