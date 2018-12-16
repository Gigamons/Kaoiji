package packets

import (
	"bytes"
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
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

	buffer.Write(helpers.GetBytes(p.PacketId))
	buffer.Write(helpers.GetBytes(byte(0)))
	buffer.Write(helpers.GetBytes(p.buffer.Len()))
	buffer.Write(p.buffer.Bytes())

	return buffer.Bytes()
}