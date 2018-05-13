package packets

import (
	"bytes"
)

// Writer is a writer.
type Writer struct {
	_buffer bytes.Buffer
}

func NewWriter() Writer {
	return Writer{}
}

func (w *Writer) Write(b []byte) {
	w._buffer.Write(b)
}

// ToByteArray returns a ByteArray
func (w *Writer) Bytes() []byte {
	return w._buffer.Bytes()
}
