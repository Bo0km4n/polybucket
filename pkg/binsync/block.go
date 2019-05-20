package binsync

import (
	"crypto/sha256"
	"io"
)

type Block struct {
	Start, End int64
	Checksum32 uint32
	Signature  [sha256.Size]byte
	HasData    bool
	RawBytes   []byte
}

var BlockSize = int64(1024)

func SetBlockSize(s int64) {
	BlockSize = s
}

func GenerateBlocks(reader io.Reader) ([]*Block, error) {
	buf := make([]byte, BlockSize)
	blocks := []*Block{}
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return blocks, err
		}
		if n == 0 {
			break
		}
	}
	return blocks, nil
}
