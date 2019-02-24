package stamp

import (
	"testing"
)

func TestBuildStamp(t *testing.T) {
	for _, tc := range []struct {
		now      string
		hostname string
		err      error
		want     string
	}{
		{"2015-10-25T10:11:12", "foo", nil, "foo@2015-10-25T10:11:12"},
		{"2015-10-25T10:11:12", "", errors.New("no name"), "<unknown>@2015-10-25T10:11:12"},
	} {
		// TODO: inject time.Now() and os.Hostname()!
		if got := BuildStamp(); tc.want != got {
			t.Errorf("BuildStamp() = %s; want = %s", got, tc.want)
		}
	}
}
