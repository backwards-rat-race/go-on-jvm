package io

import (
	"io"
)

type Serialisable interface {
	Write(w io.Writer) error
}

func ToPaddedBytes(seq, bytes uint) []byte {
	buf := make([]byte, bytes)
	for i := len(buf) - 1; seq != 0; i-- {
		buf[i] = byte(seq & 0xff)
		seq >>= 8
	}
	return buf
}

func WritePaddedBytes(w io.Writer, seq, bytes uint) error {
	_, err := w.Write(ToPaddedBytes(seq, bytes))
	return err
}

func WritePaddedBytesI(w io.Writer, seq int, bytes uint) error {
	return WritePaddedBytes(w, uint(seq), bytes)
}

func AppendPaddedBytes(b []byte, seq, bytes uint) []byte {
	return append(b, ToPaddedBytes(seq, bytes)...)
}
