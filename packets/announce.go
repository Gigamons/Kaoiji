package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

// Announce send an Yellow message to client
func (w *Writer) Announce(s string) {
	p := NewPacket(constants.BanchoAnnounce)
	p.SetPacketData(helpers.BString(s))
	w.Write(p.ToByteArray())
}
