package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

func (w *Writer) DeleteMessages(userid uint32) {
	p := constants.NewPacket(constants.BanchoUserSilenced)
	p.SetPacketData(osubinary.UInt32(userid))
	w.Write(p)
}

func (w *Writer) Silence(timeout uint32) {
	if timeout < 0 {
		timeout = 0
	}
	p := constants.NewPacket(constants.BanchoBanInfo)
	p.SetPacketData(osubinary.UInt32(timeout))
	w.Write(p)
}
