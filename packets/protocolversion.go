package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// ProtocolVersion current ProtocolVersion
func (w *Writer) ProtocolVersion(i int32) {
	p := NewPacket(constants.BanchoProtocolNegotiation)
	p.SetPacketData(osubinary.Int32(i))
	w.Write(p.ToByteArray())
}
