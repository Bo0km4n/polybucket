package binsync

import (
	"bytes"
	"io"
)

type Object struct {
	BlockSize int64
	Blocks    []*Block
}

func (o *Object) GetBlock(index int64) *Block {
	if int64(len(o.Blocks)) <= index {
		return nil
	}
	return o.Blocks[index]
}

func (o *Object) Merge(in *bytes.Reader, out *bytes.Buffer) error {
	for i := range o.Blocks {
		block := o.Blocks[i]
		if block.HasData {
			if _, err := out.Write(block.Content.Raw); err != nil {
				return err
			}
		} else {
			secReader := io.NewSectionReader(in, block.Start, block.End-block.Start)
			chunk := make([]byte, int(block.End-block.Start))
			if _, err := secReader.ReadAt(chunk, 0); err != nil {
				return err
			}
			if _, err := out.Write(chunk); err != nil {
				return err
			}
		}
	}
	return nil
}
