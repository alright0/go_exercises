package learn_io

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestLimitedReader(t *testing.T) {
	limit := 10
	lenBuffer := 5

	sReader := strings.NewReader("Hello World!")
	reader := LimitedReader{sReader, limit}
	buffer := make([]byte, lenBuffer)

	n, err := reader.Read(buffer)
	if err != nil {
		t.Errorf("First Read Fails: %v", err.Error())
	}
	if n != lenBuffer {
		t.Errorf("First Read Fails: returned bytes: %d != target limit: %d", n, limit)
	}
	fmt.Println(n, buffer, err)

	n, err = reader.Read(buffer)
	if err != nil && err != io.EOF {
		t.Errorf("Second Read Fails: %v", err.Error())
	}
	if n != lenBuffer {
		t.Errorf("First Read Fails: returned bytes: %d != target limit: %d", n, limit)
	}
	fmt.Println(n, buffer, err)
}
