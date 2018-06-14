package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// Announce send an Yellow message to client
func (w *Writer) Announce(s string) {
	p := NewPacket(constants.BanchoAnnounce)
	p.SetPacketData(osubinary.BString(s))
	w.Write(p.ToByteArray())
}
