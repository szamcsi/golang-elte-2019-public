package stamp

import (
	"fmt"
	"os"
	"time"
)

const timeFormat = "2006-01-02T15:04:05"

// BuildStamp returns a unique build stamp string.
func BuildStamp() string {
	hostname, err := os.Hostname() // TODO: refactor for testing
	if err != nil {
		hostname = "<unknown>"
	}
	now := time.Now() // TODO: refactor for testing
	return fmt.Sprintf("%s@%s", hostname, now.Format(timeFormat))
}
