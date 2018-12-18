package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Shared/shelpers"
)

func (pw *PacketWriter) Announce(message string) {
	p := new(Packet)
	p.PacketId = consts.ServerAnnounce

	shelpers.WriteBytes(&p.buffer, message, true)

	pw.WritePacket(p)
}