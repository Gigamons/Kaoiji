package packets

import (
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
)

func (pw *PacketWriter) LoginReply(reply consts.LoginReply) {
	p := new(Packet)
	p.PacketId = consts.ServerLoginReply
	p.WriteData(helpers.GetBytes(reply))
	pw.WritePacket(p)
}
