package objects

import (
	"github.com/google/uuid"
	"sync"
)

var presences []Presence
var mutPresence = sync.Mutex{}

type Presence struct {
	Token string
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
