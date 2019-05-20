package binsync

import (
	"bytes"
	"testing"
)

func TestSimpleRead(t *testing.T) {
	in := []byte{
		0x00, 0x00, 0x00, 0x01,
	}
	reader := bytes.NewReader(in)
	_, err := GenerateBlocks(reader)
	if err != nil {
		t.Fatal(err)
	}
}
