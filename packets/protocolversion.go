package packets

import (
	"git.gigamons.de/Gigamons/Kaoiji/constants"
)

// ProtocolVersion current ProtocolVersion
func (w *Writer) ProtocolVersion(i int32) {
	p := NewPacket(constants.BanchoProtocolNegotiation)
	p.SetPacketData(Int32(i))
	w.Write(p.ToByteArray())
}
