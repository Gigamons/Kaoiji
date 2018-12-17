package packets

import (
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
)

func (pw *PacketWriter) Announce(message string) {
	p := new(Packet)
	p.PacketId = consts.ServerAnnounce

	p.WriteData(helpers.GetBytes(message, true))

	pw.WritePacket(p)
}