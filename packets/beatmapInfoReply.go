package packets

import (
	"git.gigamons.de/Gigamons/Kaoiji/constants"
)

func (w *Writer) BeatmapInfoReply() {
	p := NewPacket(constants.BanchoBeatmapInfoReply)
	b := constants.BeatmapInfo{}
	p.SetPacketData(MarshalBinary(b))
	w.Write(p.ToByteArray())
}
