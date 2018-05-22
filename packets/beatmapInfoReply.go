package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

func (w *Writer) BeatmapInfoReply() {
	p := NewPacket(constants.BanchoBeatmapInfoReply)
	b := constants.BeatmapInfo{}
	p.SetPacketData(helpers.MarshalBinary(b))
	w.Write(p.ToByteArray())
}
