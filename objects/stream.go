package objects

import (
	"sync"
)

type Stream struct {
	StreamName string
	StreamTemp bool
	Tokens     []*Token
	streamLock sync.Mutex
}

var streamLock sync.Mutex
var STREAMS []*Stream

func init() {
	NewStream("main", false)
	NewStream("lobby", false)
}

// AddUser add a User to stream.
func (s *Stream) AddUser(t *Token) {
	s.streamLock.Lock()
	s.Tokens = append(s.Tokens, t)
	s.streamLock.Unlock()
}

// RemoveUser Remove user from Stream.
func (s *Stream) RemoveUser(t *Token) {
	s.streamLock.Lock()
	for i := 0; i < len(s.Tokens); i++ {
		if s.Tokens[i] == t {
			copy(s.Tokens[i:], s.Tokens[i+1:])
			s.Tokens[len(s.Tokens)-1] = nil
			s.Tokens = s.Tokens[:len(s.Tokens)-1]
		}
	}
	s.streamLock.Unlock()
}

func (s *Stream) Broadcast(b []byte, t *Token) {
	s.streamLock.Lock()
	for i := 0; i < len(s.Tokens); i++ {
		if s.Tokens[i] == t {
			continue
		}
		s.Tokens[i].Write(b)
	}
	s.streamLock.Unlock()
}

func GetStream(streamname string) *Stream {
	for i := 0; i < len(STREAMS); i++ {
		if STREAMS[i].StreamName == streamname {
			return STREAMS[i]
		}
	}
	return nil
}

func NewStream(streamname string, temp bool) *Stream {
	s := &Stream{StreamName: streamname, StreamTemp: temp}
	streamLock.Lock()
	STREAMS = append(STREAMS, s)
	streamLock.Unlock()
	return s
}

func DeleteStream(streamname string) {
	streamLock.Lock()
	for i := 0; i < len(STREAMS); i++ {
		if STREAMS[i].StreamName == streamname {
			STREAMS[i] = nil
		}
	}
	streamLock.Unlock()
}
