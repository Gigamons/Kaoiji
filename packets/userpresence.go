package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Mempler/osubinary"
)

// UserPresence Sends a User
func (w *Writer) UserPresence(t *objects.Token) {
	p := constants.NewPacket(constants.BanchoUserPresence)
	if t == nil {
		return
	}
	a := constants.UserPresenceStruct{t.User.ID, t.User.UserName, t.Status.Info.TimeZone, t.Status.Info.CountryID, t.Status.Info.Permissions, t.Status.Info.Lon, t.Status.Info.Lat, t.Status.Info.Rank}
	p.SetPacketData(osubinary.Marshal(a))
	w.Write(p)
}
