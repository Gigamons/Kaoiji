package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// ProtocolVersion current ProtocolVersion
func (w *Writer) ProtocolVersion(i uint32) {
	p := constants.NewPacket(constants.BanchoProtocolNegotiation)
	p.SetPacketData(osubinary.UInt32(i))
	w.Write(p)
}
