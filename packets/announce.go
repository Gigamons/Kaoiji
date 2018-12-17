package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
)

func (pw *PacketWriter) Announce(message string) {
	p := new(Packet)
	p.PacketId = consts.ServerAnnounce

	p.WriteData(helpers.GetBytes(message, true))

	pw.WritePacket(p)
}