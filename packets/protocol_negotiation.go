package packets

import (
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
)

func (pw *PacketWriter) ProtocolNegotiation(negotation int32) {
	p := new(Packet)
	p.PacketId = consts.ServerProtocolNegotiation

	p.WriteData(helpers.GetBytes(negotation, true))

	pw.WritePacket(p)
}