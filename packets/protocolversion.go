package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

// ProtocolVersion current ProtocolVersion
func (w *Writer) ProtocolVersion(i int32) {
	p := NewPacket(constants.BanchoProtocolNegotiation)
	p.SetPacketData(helpers.Int32(i))
	w.Write(p.ToByteArray())
}
