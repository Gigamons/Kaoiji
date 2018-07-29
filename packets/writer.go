package packets

import (
	"bytes"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
)

// Writer is a writer.
type Writer struct {
	_buffer bytes.Buffer
	_token  *objects.Token
}

func NewWriter(t *objects.Token) Writer {
	return Writer{_token: t}
}

func (w *Writer) SetToken(t *objects.Token) {
	w._token = t
}

func (w *Writer) Write(b []byte) {
	w._buffer.Write(b)
}

func (w *Writer) WritePacket(b *constants.Packet) {
	w._buffer.Write(b.ToByteArray())
}

// ToByteArray returns a ByteArray
func (w *Writer) Bytes() []byte {
	return w._buffer.Bytes()
}
