package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// LoginPermissions sends the "Client" permissions to the client.
func (w *Writer) LoginPermissions() {
	p := constants.NewPacket(constants.BanchoLoginPermissions)
	p.SetPacketData(osubinary.Int32(int32(w._token.Status.Info.ClientPerm))) // Why does client use Int32 instead of just a simple uint8 ? i mean it's not even greater then 256 bytes.
	w.Write(p)
}
