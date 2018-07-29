package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// Announce send an Yellow message to client
func (w *Writer) Announce(s string) {
	p := constants.NewPacket(constants.BanchoAnnounce)
	p.SetPacketData(osubinary.BString(s))
	w.WritePacket(p)
}
