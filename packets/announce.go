package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
)

// Announce send an Yellow message to client
func (w *Writer) Announce(s string) {
	p := NewPacket(constants.BanchoAnnounce)
	p.SetPacketData(BString(s))
	w.Write(p.ToByteArray())
}
