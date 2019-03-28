package linesservice

//go:generate protoc -I../../proto/lines --go_out=plugins=grpc:../../proto/lines lines.proto
import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/lines"
)

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		lines []string
		resp  *pb.CountResponse
		code  codes.Code
	}{
		{
			lines: []string{"1"},
			// TODO: resp
		},
		// ...
		// END SLIDE1 OMIT
		{
			lines: []string{"1", "1", "1"},
			// TODO: resp
		},
		{
			lines: []string{"1", "22"},
			// TODO: resp
		},
		{
			lines: []string{},
			// TODO: status
		},
	} {
		s := New()
		// START SLIDE2 OMIT
		resp, err := s.Count(nil, &pb.CountRequest{Lines: tc.lines})
		if got := grpc.Code(err); got != tc.code { // HL
			t.Errorf("Count(%v) err = %s; want = %s", tc.lines, got, tc.code)
			continue
		}
		// END SLIDE2 OMIT
		if tc.code == codes.OK && true { // TODO: compare resp and tc.resp
			t.Errorf("Count(%v) response differs from expected (got -> want):\n%s", tc.lines, "") // TODO: diff
		}
	}
}
