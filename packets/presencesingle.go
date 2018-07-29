package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

func (w *Writer) PresenceSingle(userid int32) {
	p := constants.NewPacket(constants.BanchoUserPresenceSingle)
	p.SetPacketData(osubinary.Int32(userid))
	w.WritePacket(p)
}
