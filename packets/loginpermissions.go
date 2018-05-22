package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
)

// LoginPermissions sends the "Client" permissions to the client.
func (w *Writer) LoginPermissions() {
	p := NewPacket(constants.BanchoLoginPermissions)
	p.SetPacketData(helpers.Int32(int32(w._token.Status.Info.ClientPerm)))
	w.Write(p.ToByteArray())
}
