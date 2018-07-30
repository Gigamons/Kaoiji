package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

func (w *Writer) DeleteMessages(userid int32) {
	p := constants.NewPacket(constants.BanchoUserSilenced)
	p.SetPacketData(osubinary.Int32(userid))
	w.WritePacket(p)
}

func (w *Writer) Silence(timeout int32) {
	if timeout < 0 {
		timeout = 0
	}
	p := constants.NewPacket(constants.BanchoBanInfo)
	p.SetPacketData(osubinary.Int32(timeout))
	w.WritePacket(p)
}
