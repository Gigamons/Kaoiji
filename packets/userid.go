package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

// UserID returns a binary encoded userid
func (w *Writer) UserID(userid int32) {
	p := NewPacket(constants.BanchoLoginReply)
	p.SetPacketData(helpers.Int32(userid))
	w.Write(p.ToByteArray())
}
