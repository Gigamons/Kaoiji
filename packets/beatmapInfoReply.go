package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

func (w *Writer) BeatmapInfoReply() {
	p := constants.NewPacket(constants.BanchoBeatmapInfoReply)
	w.Write(p)
}
