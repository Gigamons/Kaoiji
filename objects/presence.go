package objects

import (
	"bytes"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/google/uuid"
	"io"
	"sync"
)

var presences []Presence
var mutPresence = sync.Mutex{}

type Presence struct {
	Token string

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

	return &presences[index]
}
