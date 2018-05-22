package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/helpers"
)

// UserPresence Sends a User
func (w *Writer) UserPresence(t *objects.Token) {
	p := NewPacket(constants.BanchoUserPresence)
	a := constants.UserPresenceStruct{UserID: int32(t.User.ID), Username: string(t.User.UserName), Timezone: int8(t.Status.Info.TimeZone), CountryID: int8(t.Status.Info.CountryID), Permissions: int8(t.Status.Info.ClientPerm), Lon: float64(t.Status.Info.Lon), Lat: float64(t.Status.Info.Lat), Rank: int32(t.Status.Info.Rank)}
	p.SetPacketData(helpers.MarshalBinary(&a))
	w.Write(p.ToByteArray())
}
