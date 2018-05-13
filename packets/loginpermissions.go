package packets

import (
	"git.gigamons.de/Gigamons/Kaoiji/constants"
	"git.gigamons.de/Gigamons/Kaoiji/objects"
)

// LoginPermissions sends the "Client" permissions to the client.
func (w *Writer) LoginPermissions(t *objects.Token) {
	p := NewPacket(constants.BanchoLoginPermissions)
	p.SetPacketData(Int32(int32(t.Status.Info.ClientPerm)))
	w.Write(p.ToByteArray())
}
