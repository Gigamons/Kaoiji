package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
)

func (w *Writer) PresenceBundle() {
	var t []int32
	for i := 0; i < len(objects.TOKENS); i++ {
		t = append(t, objects.TOKENS[i].User.ID)
	}
	p := NewPacket(constants.BanchoUserPresenceBundle)
	p.SetPacketData(IntArray(t))
	w.Write(p.ToByteArray())
}
