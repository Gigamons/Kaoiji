package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
)

// LoginPermissions sends the "Client" permissions to the client.
func (w *Writer) LoginPermissions() {
	p := NewPacket(constants.BanchoLoginPermissions)
	p.SetPacketData(Int32(int32(w._token.Status.Info.ClientPerm)))
	w.Write(p.ToByteArray())
}
