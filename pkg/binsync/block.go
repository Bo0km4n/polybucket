package binsync

import (
	"bytes"
	"crypto/sha256"
	"hash/adler32"
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

func GenerateBlocks(src io.Reader, dst io.Reader) ([]*Block, error) {
	buf := make([]byte, BlockSize)
	srcBlocks := []*Block{}
	dstBlocks := []*Block{}
	srcStart := 0
	dstStart := 0

	// Make source object's list of signature
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n == 0 {
			break
		}
		b := &Block{
			Start:      int64(srcStart),
			End:        int64(srcStart + n),
			Signature:  sha256.Sum256(buf[:n]),
			Checksum32: adler32.Checksum(buf[:n]),
			HasData:    false,
		}
		srcBlocks = append(srcBlocks, b)
		srcStart += n
	}

	// Make blocks then compare the signature, if different from source's signature
	// build block inserted different byte array.
	srcIdx := 0
	srcBlock := &Block{}
	buf = make([]byte, BlockSize)
	for {
		if len(srcBlocks) <= srcIdx {
			srcBlock = &Block{}
		} else {
			srcBlock = srcBlocks[srcIdx]
		}

		n, err := dst.Read(buf)
		if err != nil && err != io.EOF {
			return dstBlocks, err
		}
		if n == 0 {
			break
		}
		dstBlock := &Block{
			Start:      int64(dstStart),
			End:        int64(dstStart + n),
			Checksum32: adler32.Checksum(buf[:n]),
			Signature:  sha256.Sum256(buf[:n]),
		}
		if !bytes.Equal(dstBlock.Signature[:], srcBlock.Signature[:]) {
			dstBlock.HasData = true
			dstBlock.RawBytes = make([]byte, n)
			copy(dstBlock.RawBytes, buf[:n])
		} else {
			dstBlock.HasData = false
		}

		dstBlocks = append(dstBlocks, dstBlock)
		srcIdx += 1
		dstStart += n
	}

	return dstBlocks, nil
}
