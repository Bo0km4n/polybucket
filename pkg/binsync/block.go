package binsync

import (
	"bytes"
	"crypto/sha256"
	"hash/adler32"
	"io"
)

type Content struct {
	IsNone bool
	Raw    []byte
}

type Block struct {
	Start, End int64
	Checksum32 uint32
	Signature  [sha256.Size]byte
	HasData    bool
	Content    *Content
}

type enumerator struct {
	idx      int
	src, dst []*Block
	a, b     *Block
}

var blockSize = int64(1024)

func setBlockSize(s int64) {
	blockSize = s
}

func GenerateBlocks(src io.Reader, dst io.Reader) ([]*Block, []*Block, error) {
	buf := make([]byte, blockSize)
	srcBlocks := []*Block{}
	dstBlocks := []*Block{}
	index := 0

	// Build source blocks
	for {
		n, err := src.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return srcBlocks, dstBlocks, err
		}

		sb := genBlock(index, n, buf[:n])
		sb.HasData = false
		srcBlocks = append(srcBlocks, sb)
		index += n
	}
	srcObject := &Object{
		Blocks: srcBlocks,
	}

	// Build destination blocks
	index = 0
	blockIndex := 0
	for {
		n, err := dst.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return srcBlocks, dstBlocks, err
		}

		db := genBlock(index, n, buf[:n])
		sb := srcObject.GetBlock(int64(blockIndex))
		if sb == nil || !bytes.Equal(sb.Signature[:], db.Signature[:]) {
			db.HasData = true
			db.Content = &Content{
				IsNone: false,
				Raw:    make([]byte, blockSize),
			}
			copy(db.Content.Raw, buf[:n])
		}
		dstBlocks = append(dstBlocks, db)
		index += n
		blockIndex += 1
	}

	return srcBlocks, dstBlocks, nil
}

func getBlockContentLength(blocks []*Block) int {
	length := 0
	for _, b := range blocks {
		length += int(b.End - b.Start)
	}
	return length
}

func genBlock(start, offset int, data []byte) *Block {
	return &Block{
		Start:      int64(start),
		End:        int64(start + offset),
		Checksum32: adler32.Checksum(data[:offset]),
		Signature:  sha256.Sum256(data[:offset]),
	}
}

func (e *enumerator) Next() bool {
	if len(e.src) > e.idx {
		e.a = e.src[e.idx]
	} else {
		e.a = nil
	}
	if len(e.dst) > e.idx {
		e.b = e.dst[e.idx]
	} else {
		e.b = nil
	}

	e.idx += 1
	if e.a == nil && e.b == nil {
		return false
	}
	return true
}

func (e *enumerator) Get() (*Block, *Block) {
	return e.a, e.b
}
