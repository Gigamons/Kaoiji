package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Mempler/osubinary"
)

func (w *Writer) PresenceBundle() {
	var t []int32
	for i := 0; i < len(objects.TOKENS); i++ {
		t = append(t, objects.TOKENS[i].User.ID)
	}
	p := constants.NewPacket(constants.BanchoUserPresenceBundle)
	p.SetPacketData(osubinary.IntArray(t))
	w.WritePacket(p)
}
