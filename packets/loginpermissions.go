package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// LoginPermissions sends the "Client" permissions to the client.
func (w *Writer) LoginPermissions() {
	p := constants.NewPacket(constants.BanchoLoginPermissions)
	p.SetPacketData(osubinary.Int32(int32(w._token.Status.Info.ClientPerm)))
	w.WritePacket(p)
}
