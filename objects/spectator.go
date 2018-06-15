package objects

import (
	"bytes"
	"io"
	"sync"

	"github.com/Mempler/osubinary"

	"github.com/Gigamons/Kaoiji/constants"
)

type SpectatorStream struct {
	HostToken    *Token
	streamLock   *sync.Mutex
	StreamTokens []*Token
}

type SpectatorFrame struct {
	Extra        int32
	ReplayFrames ReplayFrame
	Action       int8
	ScoreFrame   ScoreFrame
}

type ReplayFrame struct {
	ButtonState int8
	Button      int8
	MouseX      float32
	MouseY      float32
	Time        int32
}

type ScoreFrame struct {
	Time         int32
	ID           int8
	Count300     uint16
	Count100     uint16
	Count50      uint16
	CountGeki    uint16
	CountKatu    uint16
	CountMiss    uint16
	TotalScore   int32
	MaxCombo     uint16
	CurrentCombo uint16
	FC           bool
	HP           uint8
	TagByte      int8
	ScoreV2      bool
}

func (s *ScoreFrame) ScoreV2F(comboPortion, bonusPortion float64) []byte {
	out := []byte{}
	if s.ScoreV2 {
		x := osubinary.Double(comboPortion)
		y := osubinary.Double(bonusPortion)
		out = append(y, x...)
	}
	return out
}

func NewSpectatorStream(t *Token) *SpectatorStream {
	s := &SpectatorStream{HostToken: t, streamLock: &sync.Mutex{}}
	t.SpectatorStream = s
	return s
}

func (s *SpectatorStream) RemoveSpectatorStream(t *Token) {
	for i := 0; i < len(s.StreamTokens); i++ {
		s.RemoveUser(s.StreamTokens[i])
	}
	s = nil
}

func (s *SpectatorStream) AddUser(t *Token) {
	if s.AlreadySpectating(t) {
		return
	}
	s.streamLock.Lock()
	s.StreamTokens = append(s.StreamTokens, t)
	s.streamLock.Unlock()
	p1 := constants.NewPacket(constants.BanchoFellowSpectatorJoined)
	p2 := constants.NewPacket(constants.BanchoSpectatorJoined)
	p1.SetPacketData(osubinary.Int32(t.User.ID))
	p2.SetPacketData(osubinary.Int32(t.User.ID))
	t.SpectatorStream = s
	s._Broadcast(p2.ToByteArray(), false, nil)
	s._Broadcast(p1.ToByteArray(), false, t)
}

func (s *SpectatorStream) RemoveUser(t *Token) {
	for i := 0; i < len(s.StreamTokens); i++ {
		if s.StreamTokens[i] == t {
			buf := bytes.NewBuffer(nil)
			p1 := constants.NewPacket(constants.BanchoFellowSpectatorLeft)
			p2 := constants.NewPacket(constants.BanchoSpectatorLeft)
			p1.SetPacketData(osubinary.Int32(s.StreamTokens[i].User.ID))
			p2.SetPacketData(osubinary.Int32(s.StreamTokens[i].User.ID))
			buf.Write(p1.ToByteArray())
			buf.Write(p2.ToByteArray())
			s.StreamTokens[i].Write(buf.Bytes())
			s.streamLock.Lock()
			copy(s.StreamTokens[i:], s.StreamTokens[i+1:])
			s.StreamTokens[len(s.StreamTokens)-1] = nil
			s.StreamTokens = s.StreamTokens[:len(s.StreamTokens)-1]
			s.streamLock.Unlock()
		}
	}
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

func (s *SpectatorStream) Broadcast(r io.Reader, frame *SpectatorFrame) {
	f := *frame
	bf := osubinary.Marshal(f)
	pack := constants.NewPacket(constants.BanchoSpectateFrames)
	pack.SetPacketData(bf)
	z := []byte{}
	if frame.ScoreFrame.ScoreV2 {
		x, _ := osubinary.RDouble(r)
		y, _ := osubinary.RDouble(r)
		z = append(append(z, pack.ToByteArray()...), frame.ScoreFrame.ScoreV2F(x, y)...)
	}
	s._Broadcast(z, true, s.HostToken)
}

func (s *SpectatorStream) _Broadcast(b []byte, isFrame bool, ignoreSelf *Token) {
	if !isFrame {
		s.HostToken.Write(b)
	}
	for i := 0; i < len(s.StreamTokens); i++ {
		if s.StreamTokens[i] == ignoreSelf {
			continue
		}
		s.StreamTokens[i].Write(b)
	}
}
