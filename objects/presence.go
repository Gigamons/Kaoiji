package objects

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

var presences []Presence
var mutPresence = sync.Mutex{}

type Presence struct {
	token string

}

func NewUserPresence() *Presence {
	pr := Presence{}
	n, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	pr.token = n.String()

	mutPresence.Lock()
	presences = append(presences, pr)
	index := len(presences)
	mutPresence.Unlock()

	return &presences[index]
}
