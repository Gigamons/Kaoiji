package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Shared/shelpers"
)

func (pw *PacketWriter) ProtocolNegotiation(negotation int32) {
	p := new(Packet)
	p.PacketId = consts.ServerProtocolNegotiation

	shelpers.WriteBytes(&p.buffer, negotation, true)

	pw.WritePacket(p)
}