package binsync

import (
	"bytes"
	"io"
	"testing"

	"github.com/k0kubun/pp"
)

func TestSimpleRead(t *testing.T) {
	cases := []struct {
		src       io.Reader
		dst       io.Reader
		blockSize int64
	}{
		{
			src:       bytes.NewReader([]byte{0x00, 0x01}),
			dst:       bytes.NewReader([]byte{0x02}),
			blockSize: 1,
		},
		{
			src:       bytes.NewReader([]byte{0x00, 0x01, 0x02}),
			dst:       bytes.NewReader([]byte{0x00, 0x01, 0x03, 0x04, 0x05}),
			blockSize: 2,
		},
		{
			src:       bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03, 0x04}),
			dst:       bytes.NewReader([]byte{0x00, 0x01, 0x03, 0x04}),
			blockSize: 3,
		},
		{
			src:       bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03, 0x04}),
			dst:       bytes.NewReader([]byte{0x00, 0x01, 0x02}),
			blockSize: 3,
		},
	}
	for _, c := range cases {
		SetBlockSize(c.blockSize)
		_, blocks, err := GenerateBlocks(c.src, c.dst)
		if err != nil {
			t.Fatal(err)
		}
		checkBlockRawBytes(t, blocks)
	}
}

func checkBlockRawBytes(t *testing.T, blocks []*Block) {
	pp.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++")
	for i := range blocks {
		pp.Println(blocks[i])
	}
}
