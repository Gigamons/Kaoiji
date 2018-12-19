package objects

import (
	"bytes"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/Gigamons/Shared/sutilities"
	"github.com/google/uuid"
	"io"
	"sync"
)

var presences []Presence
var mutPresence = sync.Mutex{}

type Presence struct {
	Token                 string
	User                  *sutilities.User
	UserStatus            *sutilities.UserStatus

	UTCOffset             uint8

	mut sync.Mutex
	buffer bytes.Buffer
}

func (pr *Presence) WriteBytes(w io.Writer) {
	pr.mut.Lock()
	w.Write(pr.buffer.Bytes())
	pr.buffer.Reset()
	pr.mut.Unlock()
}
func (pr *Presence) WritePW(pw *packets.PacketWriter) {
	pr.mut.Lock()
	pr.buffer.Write(pw.GetBytes())
	pr.mut.Unlock()
}

func (pr *Presence) GetUserPresence() packets.UserPresence {
	p := packets.UserPresence{
		UserId:    pr.User.Id,
		UserName:  pr.User.UserName,
		Lon:       pr.UserStatus.Lon,
		Lat:       pr.UserStatus.Lat,
		CountryId: pr.UserStatus.Country,
		Timezone:  pr.UTCOffset,
	}

	//if (pr.User.Privileges & sconsts.Privilege)


	// TODO: add MrMoreGamerino.tv/shop/kissen.

	return p
}

func NewPresence() Presence {
	pr := Presence{}
	n, _ := uuid.NewRandom()
	pr.Token = n.String()
	return pr
}
func AppendPresence(pr Presence) *Presence {
	mutPresence.Lock()
	presences = append(presences, pr)
	index := len(presences)
	mutPresence.Unlock()

	return &presences[index - 1]
}
