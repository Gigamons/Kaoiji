package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Mempler/osubinary"
)

// UserPresence Sends a User
func (w *Writer) UserPresence(t *objects.Token) {
	p := NewPacket(constants.BanchoUserPresence)
	a := constants.UserPresenceStruct{int32(t.User.ID), string(t.User.UserName), int8(t.Status.Info.TimeZone), int8(t.Status.Info.CountryID), int8(t.Status.Info.ClientPerm), float64(t.Status.Info.Lon), float64(t.Status.Info.Lat), int32(t.Status.Info.Rank)}
	p.SetPacketData(osubinary.Marshal(a))
	w.Write(p.ToByteArray())
}
