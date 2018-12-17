package packets

import (
	"bytes"
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
	"io"
)

type Packet struct {
	PacketId consts.PacketId
	buffer bytes.Buffer
}

func (p *Packet) WriteData(data []byte) {
	p.buffer.Write(data)
}

func (p *Packet) GetBytes() []byte {
	buffer := new(bytes.Buffer)

	buffer.Write(helpers.GetBytes(uint16(p.PacketId)))
	buffer.Write(helpers.GetBytes(uint8(0)))
	buffer.Write(helpers.GetBytes(int32(p.buffer.Len())))
	buffer.Write(p.buffer.Bytes())

	return buffer.Bytes()
}

func (p *Packet) WriteBytes(w io.Writer) (err error) {
	if _, err = w.Write(helpers.GetBytes(uint16(p.PacketId))); err != nil {
		return
	}
	if _, err = w.Write(helpers.GetBytes(uint8(0))); err != nil {
		return
	}
	if _, err = w.Write(helpers.GetBytes(int32(p.buffer.Len()))); err != nil {
		return
	}
	if _, err = w.Write(p.buffer.Bytes()); err != nil {
		return
	}

	return
}
