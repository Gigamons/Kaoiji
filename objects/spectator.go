package objects

import (
	"bytes"
	"encoding/binary"
	"sync"

	"github.com/Gigamons/common/helpers"

	"github.com/Gigamons/Kaoiji/constants"
)

type SpectatorStream struct {
	HostToken    *Token
	streamLock   sync.Mutex
	StreamTokens []*Token
}

type spectatorFrame struct {
	Extra        int32
	ReplayFrames replayFrame
	Action       int8
	ScoreFrame   scoreFrame
}

type replayFrame struct {
	ButtonState int8
	Button      int8
	MouseX      float64
	MouseY      float64
	Time        int32
}

type scoreFrame struct {
	Time          int32
	ID            int8
	Count300      uint16
	Count100      uint16
	Count50       uint16
	CountGeki     uint16
	CountKatu     uint16
	CountMiss     uint16
	TotalScore    int32
	MaxCombo      uint16
	CurrentCombo  uint16
	FC            bool
	HP            uint8
	TagByte       int8
	_ScoreV2      bool
	_ComboPortion float32
	_BonusPortion float32
}

func (s *scoreFrame) ScoreV2() []byte {
	b := bytes.NewBuffer(nil)
	if s._ScoreV2 {
		binary.Write(b, binary.LittleEndian, int8(1))
		binary.Write(b, binary.LittleEndian, s._ComboPortion)
		binary.Write(b, binary.LittleEndian, s._BonusPortion)
	} else {
		binary.Write(b, binary.LittleEndian, int8(0))
	}
	return b.Bytes()
}

var SSTREAMS []*SpectatorStream
var sStreamLock sync.Mutex

func NewSpectatorStream(t *Token) *SpectatorStream {
	s := &SpectatorStream{HostToken: t}
	sStreamLock.Lock()
	SSTREAMS = append(SSTREAMS, s)
	sStreamLock.Unlock()
	return s
}

func RemoveSpectatorStream(t *Token) {
	sStreamLock.Lock()
	for i := 0; i < len(SSTREAMS); i++ {
		if SSTREAMS[i].HostToken == t {
			sStreamLock.Unlock()
			SSTREAMS[i] = nil
			return
		}
	}
	sStreamLock.Unlock()
}

func (s *SpectatorStream) AddUser(t *Token) {
	s.streamLock.Lock()
	s.StreamTokens = append(s.StreamTokens, t)

	b := new(bytes.Buffer)

	u := helpers.Int32(t.User.ID)

	if s.AlreadySpectating(t) {
		return
	}

	binary.Write(b, binary.LittleEndian, constants.BanchoFellowSpectatorJoined)
	binary.Write(b, binary.LittleEndian, int8(0))
	binary.Write(b, binary.LittleEndian, len(u))
	binary.Write(b, binary.LittleEndian, u)

	binary.Write(b, binary.LittleEndian, constants.BanchoSpectatorJoined)
	binary.Write(b, binary.LittleEndian, int8(0))
	binary.Write(b, binary.LittleEndian, len(u))
	binary.Write(b, binary.LittleEndian, u)

	s._Broadcast(b.Bytes())

	s.StreamTokens = append(s.StreamTokens, t)

	s.streamLock.Unlock()
}

func (s *SpectatorStream) AlreadySpectating(t *Token) bool {
	s.streamLock.Lock()
	for i := 0; i < len(s.StreamTokens); i++ {
		if s.StreamTokens[i] == t {
			s.streamLock.Unlock()
			return true
		}
	}
	s.streamLock.Unlock()
	return false
}

func (s *SpectatorStream) Broadcast(frame *spectatorFrame) {
	b := new(bytes.Buffer)
	bf := helpers.MarshalBinary(frame)

	binary.Write(b, binary.LittleEndian, constants.BanchoSpectateFrames)
	binary.Write(b, binary.LittleEndian, int8(0))
	binary.Write(b, binary.LittleEndian, len(bf))
	binary.Write(b, binary.LittleEndian, bf)
}

func (s *SpectatorStream) _Broadcast(b []byte) {
	s.streamLock.Lock()
	s.HostToken.Write(b)
	for i := 0; i < len(s.StreamTokens); i++ {
		s.StreamTokens[i].Write(b)
	}
	s.streamLock.Unlock()
}
