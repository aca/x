package fsutil

import "testing"

func TestIsSymLink(t *testing.T) {
	t.Log(IsSymLink("fs"))
}
