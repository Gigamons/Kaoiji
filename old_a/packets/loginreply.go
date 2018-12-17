package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// LoginReply returns a binary encoded Login Reply
func (w *Writer) LoginReply(reply int32) {
	p := constants.NewPacket(constants.BanchoLoginReply)
	p.SetPacketData(osubinary.Int32(reply))
	w.Write(p)
}
