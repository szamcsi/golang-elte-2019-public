// Package lines deals with lines in files.
package lines

//go:generate protoc -I../../proto/texts --go_out=plugins=grpc:../../proto/texts texts.proto
import (
	"bufio"
	"context"
	"os"

	pb "github.com/szamcsi/golang-elte-2019-public/grpc-testing/proto/texts"
)

func load(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// MinMaxCount represents line statistics information.
type MinMaxCount struct {
	Min, Max, Count int32
}

// Count counts the lines in a file and determines the minimal and maximal line length.
// It returns an error, if the file was not readable.
func Count(ctx context.Context, client pb.TextsServiceClient, path string) (*MinMaxCount, error) {
	lines, err := load(path)
	if err != nil {
		return nil, err
	}
	resp, err := client.Count(ctx, &pb.CountRequest{Lines: lines})
	if err != nil {
		return nil, err
	}
	return &MinMaxCount{
		Min:   resp.Min,
		Max:   resp.Max,
		Count: resp.Count,
	}, nil
}
