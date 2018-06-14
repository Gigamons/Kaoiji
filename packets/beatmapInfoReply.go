package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

func (w *Writer) BeatmapInfoReply() {
	p := NewPacket(constants.BanchoBeatmapInfoReply)
	b := constants.BeatmapInfo{}
	p.SetPacketData(osubinary.Marshal(b))
	w.Write(p.ToByteArray())
}
