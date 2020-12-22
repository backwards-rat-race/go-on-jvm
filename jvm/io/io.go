package io

import "io"

func ToPaddedBytes(seq int, bytes int) []byte {
	buf := make([]byte, bytes)
	for i := len(buf) - 1; seq != 0; i-- {
		buf[i] = byte(seq & 0xff)
		seq >>= 8
	}
	return buf
}

func WritePaddedBytes(w io.Writer, seq int, bytes int) error {
	_, err := w.Write(ToPaddedBytes(seq, bytes))
	return err
}
