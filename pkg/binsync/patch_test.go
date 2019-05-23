package binsync

import (
	"bytes"
	"testing"

	"github.com/k0kubun/pp"
)

func TestOnePatch(t *testing.T) {

	cases := []struct {
		in       []byte
		expected []byte
		bsize    int64
	}{
		{
			in:       []byte{0x01, 0x01, 0x01},
			expected: []byte{0x01, 0x02, 0x03},
			bsize:    3,
		},
		{
			in:       []byte{0x01, 0x01, 0x01, 0x02, 0x03},
			expected: []byte{0x01, 0x01, 0x01, 0x02, 0x03},
			bsize:    3,
		},
		{
			in:       []byte{0x01, 0x01, 0x01},
			expected: []byte{0x01, 0x01, 0x01, 0x02, 0x03},
			bsize:    3,
		},
		{
			in:       []byte{0x01, 0x01, 0x01, 0x02, 0x03},
			expected: []byte{0x01, 0x01, 0x02},
			bsize:    3,
		},
	}

	for _, c := range cases {
		src := bytes.NewReader(c.in)
		dst := bytes.NewReader(c.expected)

		setBlockSize(c.bsize)

		_, dstBlocks, err := GenerateBlocks(src, dst)
		if err != nil {
			t.Fatal(err)
		}

		object := &Object{
			BlockSize: blockSize,
			Blocks:    dstBlocks,
		}

		// Reset src io
		newSrc := bytes.NewReader(c.in)
		var buffer bytes.Buffer

		if err := object.Merge(newSrc, &buffer); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(buffer.Bytes()[:len(c.expected)], c.expected) {
			pp.Println(buffer.Bytes()[:len(c.expected)], c.expected)
			t.Error("Not matched buffer.Bytes() and expected")
		}
	}
}
