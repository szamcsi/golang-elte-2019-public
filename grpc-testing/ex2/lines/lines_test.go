package lines

import (
	"context"
	"github.com/szamcsi/golang-elte-2019-public/grpc-testing/ex1/linesservice"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"path/filepath"
	"testing"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/lines"
)

const (
	testDir = "../../testdata/"
	port    = ":54321"
)

func server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLinesServiceServer(s, linesservice.New())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go server()
	os.Exit(m.Run())
}

// OMITTED: func TestLoad(t *testing.T)

// START OMIT
func TestCount(t *testing.T) {
	// TODO: create a 'client', similar to the lc.go:main() function!
	// END OMIT

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
		mmc, err := Count(context.Background(), client, filepath.Join(testDir, tc.path))
		if tc.wantErr != (err != nil) {
			t.Errorf("Count(%q) err != nil is %v; want %v", tc.path, err != nil, tc.wantErr)
			continue
		}
		if !cmp.Equal(mmc, tc.mmc) {
			t.Errorf("Count(%q) = %v, _; want = %v, _", tc.path, mmc, tc.mmc)
		}
	}
}
