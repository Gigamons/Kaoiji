package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
)

func (pw *PacketWriter) ProtocolNegotiation(negotation int32) {
	p := new(Packet)
	p.PacketId = consts.ServerProtocolNegotiation

	helpers.WriteBytes(&p.buffer, negotation, true)

	pw.WritePacket(p)
}