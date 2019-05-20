package binsync

import (
	"bytes"
	"io"
	"testing"

	"github.com/k0kubun/pp"
)

func TestSimpleRead(t *testing.T) {
	cases := []struct {
		src io.Reader
		dst io.Reader
	}{
		{
			src: bytes.NewReader([]byte{0x00, 0x01}),
			dst: bytes.NewReader([]byte{0x02}),
		},
	}
	SetBlockSize(1)
	for _, c := range cases {
		_, blocks, err := GenerateBlocks(c.src, c.dst)
		if err != nil {
			t.Fatal(err)
		}
		checkBlockRawBytes(t, blocks)
	}
}

func checkBlockRawBytes(t *testing.T, blocks []*Block) {
	for i := range blocks {
		pp.Println(blocks[i].RawBytes)
	}
}
