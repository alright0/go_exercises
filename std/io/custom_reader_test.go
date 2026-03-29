package learn_io

import (
	"testing"
)

func TestCustomReader(t *testing.T) {
	LenBuff := 8
	result := "AAAAAAAA"

	p := make([]byte, LenBuff)
	reader := &MyReader{}
	bytes, err := reader.Read(p)

	if err != nil {
		t.Errorf("%v", err.Error())
	}

	if bytes != LenBuff {
		t.Errorf("bytes should be %d", LenBuff)
	}

	if string(p) != result {
		t.Errorf("string should be %s, got: %s", result, string(p))
	}
}
