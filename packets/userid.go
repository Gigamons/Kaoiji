package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// UserID returns a binary encoded userid
func (w *Writer) UserID(userid int32) {
	p := constants.NewPacket(constants.BanchoLoginReply)
	p.SetPacketData(osubinary.Int32(userid))
	w.WritePacket(p)
}
