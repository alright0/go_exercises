package learn_io

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestTeeReader(t *testing.T) {
	str := "Hello World!"

	var log bytes.Buffer

	sReader := strings.NewReader(str)
	r := TeeReader{sReader, &log}
	data, _ := io.ReadAll(&r)

	if string(data) != str {
		t.Errorf("Stream from reader to log FAILED: expected: %s, got: %s", str, string(data))
	}
}
