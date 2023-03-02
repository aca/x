package fsutil

import "testing"

func TestWriteFile(t *testing.T) {
    WriteFile("wer", []byte("hello world"))
}
