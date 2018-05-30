package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

func (w *Writer) PresenceSingle(userid int32) {
	p := NewPacket(constants.BanchoUserPresenceSingle)
	p.SetPacketData(helpers.Int32(userid))
	w.Write(p.ToByteArray())
}
