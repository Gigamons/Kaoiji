package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

func (w *Writer) PresenceSingle(userid uint32) {
	p := constants.NewPacket(constants.BanchoUserPresenceSingle)
	p.SetPacketData(osubinary.UInt32(userid))
	w.Write(p)
}
